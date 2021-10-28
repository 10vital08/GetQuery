package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var query string

func handleQuery(w http.ResponseWriter, r *http.Request) {
	query := string(r.URL.Query().Get("query"))
	fmt.Fprintf(w, query)
}

func handleHtmlPage(w http.ResponseWriter, r *http.Request) {
	content := `<html lang="ru">
				<head>
					<meta charset="UTF-8">
					<title>Title</title>
				</head>
				<body>
					<h1>Привет, Алексей!</h1>
					<ul>
						<li>Это третье задание</li>
					</ul>
				</body>
				</html>`

	//html-шаблон
	tmpl, err := template.New("example").Parse(content)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	tmpl.Execute(w, content)
}

func main() {
	webServer := http.NewServeMux()
	webServer.HandleFunc("/", handleQuery) // установка роутера
	webServer.HandleFunc("/html", handleHtmlPage)
	log.Println("html-страница на http://127.0.0.1:8080/html")
	err := http.ListenAndServe(":8080", webServer) // установка порта
	log.Fatal(err)

}
