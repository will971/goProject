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
	r.HandleFunc("/qui-utilise-go", goUsersHandler).Methods("GET")
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.HandleFunc("/post", postHandler).Methods("GET")
	r.HandleFunc("/creators/Robert-Griesmer", robertGriesmerHandler).Methods("GET")
	r.HandleFunc("/creators/Rob-Pike", robPikeHandler).Methods("GET")
	r.HandleFunc("/creators/Ken-Thompson", kenThompsonHandler).Methods("GET")
	r.HandleFunc("/go/code", goCodeHandler).Methods("GET")
	r.HandleFunc("/go/routines", goRoutinesHandler).Methods("GET")
	r.HandleFunc("/pourquoi-choisir-go", whyGoHandler).Methods("GET")
	http.Handle("/", r)
	http.Handle("/godBy", r)
	http.ListenAndServe(":8080", nil)

}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}
func goRoutinesHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "goRoutines.html", nil)
}
func goCodeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "goCode.html", nil)
}
func goUsersHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "goUsers.html", nil)
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contact.html", nil)
}
func whyGoHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "whyGo.html", nil)
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "single-post.html", nil)
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
