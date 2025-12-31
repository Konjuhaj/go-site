package main

import (
	"net/http"
	"time"
	"fmt"
	"log"
	"html"
)

func main() {
	http.HandleFunc("/", showInfo)

	files := http.FileServer(http.Dir("/var/www"))

	http.Handle("/about/", http.StripPrefix("/about/", files))

	err := http.ListenAndServe(":8999", nil)

	if err != nil {
		log.Fatal("Listen And Serve; ", err) 
	}
}

func showInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current time: ", time.Now())

	fmt.Fprintln(w, "Current path: ", html.EscapeString(r.URL.Path))
}


