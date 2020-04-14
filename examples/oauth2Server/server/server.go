package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("auth is %v", r)

		_ = r.ParseForm()
		password := r.Form.Get("password")
		fmt.Printf("password is %v", password)
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("token is %v", r)
		_ = r.ParseForm()
		password := r.Form.Get("password")
		fmt.Printf("password is %v", password)

	})
	log.Fatal(http.ListenAndServe(":9096", nil))
}
