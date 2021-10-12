package main

import (
	"fmt"
	"log"
	"net/http"
)

var query string

func makeQuery(w http.ResponseWriter, r *http.Request, query string)  {
	fmt.Fprintf(w, "Query")
	w.Write([]byte(q))
}

func returnAnswer()  {
	
}

func main(){
	query = "123"

	vebServer := http.NewServeMux()
	http.HandleFunc("/", makeQuery)//установка роутера
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", vebServer)//установка порта
	log.Fatal(err)

	resp, err := http.Get("http://127.0.0.1:8080", query)
	defer resp.Body.Close()

}
