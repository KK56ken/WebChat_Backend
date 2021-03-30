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

    body := r.Body
    defer body.Close()

    fmt.Printlen(r.Body)

    var users []model.User
    if err := json.Unmarshal(body, &users); err != nil{
        log.Fatal(err)
    }
    for _, p := range persons{
        fmt.Printf("%d : %s\n",p.Email, p.Name, p.Password)
    }

    fmt.Fprintf(w, "OK")
}