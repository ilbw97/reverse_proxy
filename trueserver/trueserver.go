package main
import (
	"io"
	"net/http"
	"log"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":2003", nil))
}