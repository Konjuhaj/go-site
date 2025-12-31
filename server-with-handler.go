 package main

 import (
	"net/http"
	"fmt"
 )

 func main() {

	http.Handle("/", middleware1(middleware2(http.HandlerFunc(finalmiddleware))))

	err := http.ListenAndServe(":8999", nil)

	if err != nil {
		fmt.Printf("error: ", err)
	}
 }

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Executing middleware1...")
		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware1 again...")
	})

}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Executing middleware2...")
		if r.URL.Path != "/" {
			return
		}

		next.ServeHTTP(w, r)
		fmt.Fprintln(w, "Executing middleware 2 again...") 
	})
}

func finalmiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Executing final...")
	fmt.Fprintln(w, "Done")
}
