package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

type information struct {
	Host   string `json:"host"`
	Origin string `json:"origin"`
	Port   int    `json:"port"`
}

// type Informations struct{
//     Informations []information `json:"informations"`
// }

var info *information

// var test *informtaion = &informtaion{
//     Host:   "www.example.com",
// 	Origin: "1.1.1.1",
// 	Port:   8080,
// }

// type reverse_proxy struct {
// 	rpserver *httputil.ReverseProxy
// }

func s_info_Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		get_info, err := json.Marshal(info)
		if err != nil {
			log.Println(err)
		}
		io.WriteString(w, string(get_info))
		w.Write([]byte("\n"))
	case "POST":
		new_info := &info
		err := json.NewDecoder(r.Body).Decode(new_info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Printf("sending %v \n", *new_info)
		info = *new_info
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "invalid type of method")
	}
}

func reverse_p_server(w http.ResponseWriter, r *http.Request) {
	log.Println("ServeHTTP")
	target_url := "http://" + info.Host + ":" + strconv.Itoa(info.Port)
	target, err := url.Parse(target_url)
	log.Println(target)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Director = func(r *http.Request) {
		r.URL.Host = info.Origin
		log.Println(r.URL.Host)
	}
	log.Printf("before proxy : %v \n", proxy)
	proxy.ServeHTTP(w, r)
	log.Printf("proxy : %v \n", proxy)
}

func main() {
	http.HandleFunc("/", s_info_Handler)
	http.ListenAndServe(":301", nil)
}

// target_url := test.Host + ":" + strconv.Itoa(test.Port)
// target, err := url.Parse(target_url)
// if err != nil {
//     panic(err)
// }
// proxy := httputil.NewSingleHostReverseProxy(target)
// proxy.Director = func(r *http.Request) {
//     r.URL.Scheme = "http"
//     r.URL.Host = test.Origin
// }
// proxy.ServeHTTP(w, r)

// func (server_info *information) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	target_url := "http://" + server_info.Host + ":" + strconv.Itoa(server_info.Port)
// 	target, err := url.Parse(target_url)
// 	if err != nil {
// 		panic(err)
// 	}
// 	proxy := httputil.NewSingleHostReverseProxy(target)
// 	proxy.Director = func(r *http.Request) {
// 		r.URL.Host = server_info.Origin
// 	}
// 	log.Printf("before proxy : %v \n", proxy)
// 	proxy.ServeHTTP(w, r)
// 	log.Printf("proxy : %v \n", proxy)
// }
// func (server_info *information) start_rpserver() {
// 	i := &server_info
// 	http.ListenAndServe(fmt.Sprintf(":%d", info.Port), *i)
// }
// func main() {
// 	http.HandleFunc("/", s_info_Handler)
// 	http.ListenAndServe(":301", nil)
// 	go info.start_rpserver()
// }
// package main

// import(
// 	"fmt"
// 	"log"
// 	"net/url"
// 	"net/http"
// 	"net/http/httputil"
// )

// // func main() {
// // 	target, err := url.Parse("http://localhost:301")
// // 	if err != nil{
// // 		panic(err)
// // 	}
// // 	proxy := httputil.NewSingleHostReverseProxy(target)
// // 	http.HandleFunc("/test",handler(proxy))
// // 	err = http.ListenAndServe(":301", nil)
// // 	if err != nil{
// // 		panic(err)
// // 	}
// // }
// // func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request){
// // 	return func(w http.ResponseWriter, r *http.Request){
// // 		log.Println(r.URL)
// // 		r.URL.Path = "/"
// // 		p.ServeHTTP(w, r)
// // 	}
// // }
