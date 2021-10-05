package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func makeQuery(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Query")
}

func returnAnswer()  {
	
}

func main(){
	http.HandleFunc("/", makeQuery)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	resp, err := http.Get(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}