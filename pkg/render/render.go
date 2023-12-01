package render

import (
	"fmt"
	"net/http"
	"text/template"
)

//Template rendering
func RenderTemplate (w http.ResponseWriter, tpml string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tpml)
	err := parsedTemplate.Execute(w,  nil)
	if err != nil {
		fmt.Println("Error occured", err)
	}
}