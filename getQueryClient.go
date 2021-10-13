package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var query string

func makeQuery(w http.ResponseWriter, r *http.Request)  {
	//fmt.Fprintf(w, "Query")
	//w.Write([]byte(query))

	responseData, err := ioutil.ReadAll(r.Body)
	if err!= nil{
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Println(responseString)
}

func main(){
	query = "test"
	//vebServer := http.NewServeMux()
	http.HandleFunc("/", makeQuery)//установка роутера
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", nil)//установка порта
	log.Fatal(err)

	resp, err := http.Get("http://127.0.0.1:8080?query=" + query)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
