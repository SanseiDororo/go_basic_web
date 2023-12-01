package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello from the server")
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})

	http.ListenAndServe(":8080", nil)
}
