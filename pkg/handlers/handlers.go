package handlers

import (
	"net/http"

	"github.com/SanseiDororo/go_basic_web/pkg/render"
)

//Home page handler
func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home-page.tpml")
}


//About page handler
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about-page.tpml")
}