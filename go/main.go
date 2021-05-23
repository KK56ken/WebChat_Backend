package main

import(
    "net/http"
    "webchat/controller"
    "log"
    // "time"
    // "fmt"
    // "os"
)

func main() {
    http.HandleFunc("/api/register", controller.Register)
    http.HandleFunc("/api/login", controller.Login)
    http.HandleFunc("/api/autoLogin", controller.AutoLogin)
    http.HandleFunc("/api/message", controller.Message)
    http.HandleFunc("/api/friend", controller.Friend)
    http.HandleFunc("/ws", controller.Connection)
    // http.HandleFunc("/private", controller.JwtMiddleware.Handler(controller.Private))
    // http.HandleFunc("/public", controller.Public)
    go controller.HandleMessages()

    // ticker := time.NewTicker(time.Minute * 1)

    // defer ticker.Stop()
    // count := 0
    // <-ticker.C
    // fmt.Println("test")
    // for {
    //     select {
    //     case <-ticker.C:
    //         fmt.Println(count)
    //         count++
    //     }
    // }

    err := http.ListenAndServe(":9000", nil)

    if err != nil{
        log.Fatal("ListenAndServe: ", err)
    }

}


