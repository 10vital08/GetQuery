package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var query string

func handleQuery(w http.ResponseWriter, r *http.Request) {
	query := string(r.URL.Query().Get("query"))
	fmt.Fprintf(w, query)
}

func handleHtmlPage(w http.ResponseWriter, r *http.Request) {
	content := filepath.Join("public", "html", "Page.html")

	//создание html-шаблона
	tmpl, err := template.ParseFiles(content)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//вывод шаблона
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func main() {
	webServer := http.NewServeMux()
	webServer.HandleFunc("/", handleQuery) // установка роутера
	webServer.HandleFunc("/html", handleHtmlPage)
	log.Println("html-страница на http://127.0.0.1:8080/html")
	err := http.ListenAndServe(":8080", webServer) // установка порта
	log.Fatal(err)

}
