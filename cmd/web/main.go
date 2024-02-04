// cmd/web/main.go

package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var (
	templates       = template.Must(template.ParseFiles("ui/html/home.html"))
	templateCreate  = template.Must(template.ParseFiles("ui/html/create.html"))
	templateContact = template.Must(template.ParseFiles("ui/html/contact.html"))
	templateFilter  = template.Must(template.ParseFiles("ui/html/filtered.html"))
	templateEdit    = template.Must(template.ParseFiles("ui/html/edit.html"))
	templateArticle = template.Must(template.ParseFiles("ui/html/blog.html"))
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", homeHandler).Methods("GET")
	rtr.HandleFunc("/create.html", createHandler).Methods("GET")
	rtr.HandleFunc("/save_article", save_blog).Methods("POST")
	rtr.HandleFunc("/contact.html", contactHandler).Methods("GET")
	rtr.HandleFunc("/filtered/{for_who}", filterHandler).Methods("GET")
	rtr.HandleFunc("/save_contact", save_contact).Methods("POST	")
	rtr.HandleFunc("/edit/{id}", editBlogHandler).Methods("GET", "POST")
	rtr.HandleFunc("/article/{id}", blogHandler).Methods("GET")
	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}
