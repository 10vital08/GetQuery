package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var query string

func handleQuery(w http.ResponseWriter, r *http.Request) {
	query := string(r.URL.Query().Get("query"))
	fmt.Fprintf(w, query)
}

func handleHtmlPage(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("public", "html", "page.html")

	//создаем html-шаблон
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//выводим шаблон клиенту в браузер
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
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", webServer) // установка порта
	log.Fatal(err)

}
