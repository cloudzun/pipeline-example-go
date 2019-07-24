package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "Hello World! Abraham was here 2019-7-24 18:00:58.69"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
