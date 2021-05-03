package main

import(
    "net/http"
    "webchat/controller"
)

func main() {
    http.HandleFunc("/api/register", controller.Register)
    http.HandleFunc("/api/login", controller.Login)
    http.HandleFunc("/api/chat", controller.Chat)
    http.HandleFunc("/api/friend", controller.Friend)

    http.ListenAndServe(":9000", nil)
}


