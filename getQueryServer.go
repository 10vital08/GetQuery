package main

import (
	"fmt"
	"log"
	"net/http"
)
	var q string

func makeQuery(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte(q))

	//fmt.Fprintf(w,"Test")
}

func main(){
	q = "test"
	http.HandleFunc("/", makeQuery)//установка роутера
	log.Println("http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", nil)//установка порта
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
	defer resp.Body.Close()

	/*for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil{
			break
		}
	}*/

}
