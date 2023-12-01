package main

import (
	"fmt"
	"net/http"

	"github.com/SanseiDororo/go_basic_web/pkg/handlers"
)

const PortNumber = ":8080"


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("The server is starting at %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
