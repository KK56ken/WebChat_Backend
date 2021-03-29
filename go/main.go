package main

import(
    "fmt"
    "net/http"
    "webchat/controller"
)

func main() {
    http.HandleFunc("/api/user", controller.Handler)
    http.HandleFunc("/api/test", testhandler)
    http.ListenAndServe(":9000", nil)
}

func testhandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "test")
}

