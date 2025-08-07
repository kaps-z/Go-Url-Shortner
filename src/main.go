package main

import (
    "fmt"
    "log"
    "net/http"
	"html/template"
)

var tmpl *template.Template

func main() {
	InitRedis()
	tmpl = template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/r/", redirectHandler)

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
