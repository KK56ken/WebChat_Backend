package controller

import(
    "fmt"
    "net/http"
    "webchat/model"
    "encoding/json"
    "log"
    // "io/ioutil"
)

func Register(w http.ResponseWriter, r *http.Request) {

    var user model.User

    body := r.Body
    defer body.Close()

    json.NewDecoder(r.Body).Decode(&user)

    fmt.Printlen(r.Body)
    fmt.Printlen(user)

    model.UserCreate(user)

    fmt.Fprintf(w, "OK")
}

// {
//    "email":"test@example.com"
//    "user":"ken"
//    "password":"password123"
// }