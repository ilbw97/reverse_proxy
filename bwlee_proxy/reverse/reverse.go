package reverse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

type rpserver_handler struct {
	// info  *information
	proxy *httputil.ReverseProxy
}
type information struct {
	Host   string `json:"host"`
	Origin string `json:"origin"`
	Port   int    `json:"port"`
}

var rpserver_info information

func (rpserver *rpserver_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rpserver.proxy.ServeHTTP(w, r)
}

func rvproxy_handle() {
	data, err := os.Open("/root/reverse_proxy/bwlee_proxy/information.json")
	if err != nil {
		log.Fatal(err)
	}
	bytevalue, _ := ioutil.ReadAll(data)

	json.Unmarshal(bytevalue, &rpserver_info)

	target_url := rpserver_info.Host + ":" + strconv.Itoa(rpserver_info.Port)
	target, err := url.Parse(target_url)
	if err != nil {
		panic(err)
	}

	log.Printf("taget_url : %v\n", target_url)
	log.Printf("taget : %v\n", target)
	log.Printf("target.Scheme : %v\n", target.Scheme)
	log.Printf("target.Host : %v\n", target.Host)
	log.Printf("target.Port : %v\n", target.Port())
	rvproxy := httputil.NewSingleHostReverseProxy(target)

	rvproxy.Director = func(r *http.Request) {
		r.Header.Add("X-Forwarded-Host", r.Host)
		r.Header.Add("X-Origin-Host", rpserver_info.Origin)
		r.URL.Scheme = "http"
		r.URL.Host = rpserver_info.Origin

		log.Printf("info.Origin : %v\n", rpserver_info.Origin)
		log.Printf("%v -> %v \n", rpserver_info.Host, r.URL.Host)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/new_server", rvproxy.ServeHTTP)
	// rvproxy := &httputil.ReverseProxy{Director: director}
	// rpserver := rpserver_handler{proxy: rvproxy}
	// http.Handle("/new_server", &rpserver)
	// // http.ListenAndServe(":301", nil)
	go log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", rpserver_info.Port), mux))
	log.Printf("rpserver port : %v\n", rpserver_info.Port)
}
