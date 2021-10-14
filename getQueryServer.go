package main

import (
	"fmt"
	"log"
	"net/http"
)

var query string

func handleQuery(w http.ResponseWriter, r *http.Request)  {
	query := string(r.URL.Query().Get("query"))
	fmt.Fprintf(w, query)
}

func main(){
	webServer := http.NewServeMux()
	webServer.HandleFunc("/", handleQuery) // установка роутера
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", webServer) // установка порта
	log.Fatal(err)
}