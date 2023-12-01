package main

import (
	"fmt"
	"net/http"
)

const PortNumber = ":8080"


func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("The server is starting at %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
