package main 

import ( 
	"net/http"
	"html"
	"log"
	"fmt"
	"time"
)
func main() {
	http.HandleFunc("/", showInfo)
	http.HandleFunc("/about", showAbout)

	err := http.ListenAndServe(":8999", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func showInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Current time is: ", time.Now())
	fmt.Fprintln(w, "URL Path: ", html.EscapeString(r.URL.Path))
}

func showAbout(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
