package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const PortNumber = ":8080"

//Home page handler
func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "home-page.tpml")
}


//About page handler
func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "about-page.tpml")
}

//Template rendering
func renderTemplate (w http.ResponseWriter, tpml string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tpml)
	err := parsedTemplate.Execute(w,  nil)
	if err != nil {
		fmt.Println("Error occured", err)
	}
}



func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("The server is starting at %s", PortNumber))
	_ = http.ListenAndServe(PortNumber, nil)
}
