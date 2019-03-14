package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("templates/*.html"))
	setHandler()
}

func setHandler() {

	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/blog", blogHandler).Methods("GET")
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.HandleFunc("/post", postHandler).Methods("GET")
	r.HandleFunc("/project", projectHandler).Methods("GET")
	r.HandleFunc("/creators/Robert-Griesmer", robertGriesmerHandler).Methods("GET")
	r.HandleFunc("/creators/Rob-Pike", robPikeHandler).Methods("GET")
	r.HandleFunc("/creators/Ken-Thompson", kenThompsonHandler).Methods("GET")

	http.Handle("/", r)
	http.Handle("/godBy", r)
	http.ListenAndServe(":8080", nil)

}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}
func blogHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "blog.html", nil)
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "about.html", nil)
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contact.html", nil)
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "single-post.html", nil)
}
func projectHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "single-project.html", nil)
}
func robPikeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "robPike.html", nil)
}
func robertGriesmerHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "robertGriesmer.html", nil)
}
func kenThompsonHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "kenThompson.html", nil)
}
