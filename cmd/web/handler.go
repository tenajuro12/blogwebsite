// handlers.go
package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	allBlogs()
	err := templates.ExecuteTemplate(w, "home.html", posts)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	err := templateCreate.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	err := templateContact.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	forWho := params["for_who"]
	filteredBlogs := filterBlogs(forWho)
	err := templateFilter.ExecuteTemplate(w, "filtered.html", filteredBlogs)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AituNews")
	if err != nil {
		panic(err)
	}
	params := mux.Vars(r)
	id := params["id"]
	var article Article
	err = db.QueryRow("SELECT * FROM articles WHERE id = ?", id).Scan(&article.Id, &article.Title, &article.Anons, &article.Full_text, &article.For_who, &article.CreatedAt)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	err = templateArticle.ExecuteTemplate(w, "blog.html", article)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
