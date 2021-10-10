package main

import (
	"fmt"
	"log"
	"net/http"
)

func makeQuery(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Query")
}

func returnAnswer()  {
	
}

func main(){
	//query := 123

	vebServer := http.NewServeMux()
	http.HandleFunc("/", makeQuery)//установка роутера
	err := http.ListenAndServe(":8080", vebServer)//установка порта
	log.Fatal(err)


	//resp, err := http.Get()
	/*resp, err2 := http.Get(":8080")
	if err2 != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil{
			break
		}
	}*/

}
