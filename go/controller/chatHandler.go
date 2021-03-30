package controller

import(
    "fmt"
    "net/http"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}