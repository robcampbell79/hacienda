package main

import(
	"fmt"
	"net/http"
	//"html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])

	fmt.Println("Got em")
}

func main() {
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}