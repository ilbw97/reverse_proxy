package main

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// 	"flag"
// )

// func main(){
// 	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// 		fmt.Fprintf(w, "this call was relayed by the reverse proxy")
// 	}))
// 	defer backendServer.Close()

// 	rpURL, err := url.Parse(backendServer.URL)
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	frontendProxy := httptest.NewServer(httputil.NewSingleHostReverseProxy(rpURL))
// 	defer frontendProxy.Close()

// 	resp, err := http.Get(frontendProxy.URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	b, err := io.ReadAll(resp.Body)
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", b)
// }

// func main() {
//     // t_url := "http://"+ test.Host + ":"+ strconv.Itoa(test.Port)
//     url, err := url.Parse("http://localhost:301")
//     if err != nil{
//         panic(err)
//     }
//     port := flag.Int("p",80,"port")
//     flag.Parse()

//     director := func(r *http.Request) {
//         r.URL.Scheme = url.Scheme
//         r.URL.Host = url.Host
//     }
//     reverseProxy := &httputil.ReverseProxy{Director: director}
//     handler := handler{proxy: reverseProxy}
//     http.Handle("/", handler)
//     http.ListenAndServe(fmt.Sprintf("%d", *port), nil)
    
// }

// type handler struct{
//     proxy *httputil.ReverseProxy
// }

// func (h handler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
//     h.proxy.ServeHTTP(w, r)
// }

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	trueServer := "http://127.0.0.1:2003"

	url, err := url.Parse(trueServer)
	if err != nil{
		log.Println(err)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":2002", nil))
}