package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type information struct {
	Host   string `json:"host"`
	Origin string `json:"origin"`
	Port   int    `json:"port"`
}

type informations = []information

var info *information

func s_info_Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		get_info, err := json.Marshal(info)
		if err != nil {
			log.Println(err)
		}
		io.WriteString(w, string(get_info))
		w.Write([]byte("\n"))
		// err = ioutil.WriteFile("info.json", get_info, 0644)
		jsonfile, err := os.Create("./information.json")
		if err != nil {
			log.Fatal(err)
		}
		defer jsonfile.Close()

		jsonfile.Write(get_info)

		jsonfile.Close()
		fmt.Println("JSON DATA WRRITEN TO \n", jsonfile.Name())

	case "POST":
		new_info := &info
		err := json.NewDecoder(r.Body).Decode(new_info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		info = *new_info

		log.Printf("sending %v \n", info)
		log.Printf("Host : %v, Origin : %v, Port : %v \n", info.Host, info.Origin, info.Port)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid type of method")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/new_info", s_info_Handler)
	http.ListenAndServe(":301", mux)
}
