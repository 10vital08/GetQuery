package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"io"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path))
}

func main() {
	http.HandleFunc("/", handler)
}
