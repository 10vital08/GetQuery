package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	//"strconv"
)

var query string

func makeQuery(w http.ResponseWriter, r *http.Request)  {
	/*id, err := strconv.Atoi(r.URL.Query().Get("query"))
	responseString := string(id)
	fmt.Fprintf(w, responseString)*/
	
	query2 := r.URL.Query().Get("query")
	responseString := string(query2)
	fmt.Fprintf(w, responseString)
}

func createQuery (w http.ResponseWriter, r *http.Request)  {
	resp, err := http.Get("http://127.0.0.1:8080?query=test" + query)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func main(){
	query = "test"
	vebServer := http.NewServeMux()
	vebServer.HandleFunc("/", makeQuery)//установка роутера
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", vebServer)//установка порта
	log.Fatal(err)

	vebServer.HandleFunc("/", createQuery)
	vebServer.HandleFunc("/", makeQuery)
}
