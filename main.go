package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
	}

	account := r.FormValue("account")
	password := r.FormValue("password")
	fmt.Fprintf(w, "account: %s\npassword: %s", account, password)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello there!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/greet", greetHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("cannot start server")
	}
}
