package main

import "net/http"

//Home page handler
func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "home-page.tpml")
}


//About page handler
func About(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "about-page.tpml")
}