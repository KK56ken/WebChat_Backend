package controller

import(
    "fmt"
    "net/http"
    "webchat/model"
    "encoding/json"
    "io/ioutil"
)

func Register(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )

    if r.Method == http.MethodPost{

        fmt.Println("POST通ったよ")

        user := model.User{}

        requestBytes, err := ioutil.ReadAll(r.Body)

        fmt.Println(requestBytes)
        if err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            fmt.Println("io error")
            return
        }

        fmt.Println("unmarshal前")
        fmt.Println(string(requestBytes))

        if err := json.Unmarshal(requestBytes, &user); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            fmt.Println("JSON Unmarshal error:", err)
            return
        }
        fmt.Printf("%+v",&user)
        fmt.Println("unmarshal後")

        model.UserCreate(user)

        fmt.Fprintf(w, "OK")
    }
}

// {
//    "email":"test@example.com"
//    "user":"ken"
//    "password":"password123"
// }