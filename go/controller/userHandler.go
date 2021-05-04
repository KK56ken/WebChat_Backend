package controller

import(
    "fmt"
    "net/http"
    "webchat/model"
    "encoding/json"
    "io/ioutil"
    "strconv"
    "reflect"
)

func Register(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )

    if r.Method == http.MethodPost{
        user := model.User{}

        requestBytes, err := ioutil.ReadAll(r.Body)

        if err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            fmt.Println("io error")
            return
        }

        if err := json.Unmarshal(requestBytes, &user); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            fmt.Println("JSON Unmarshal error:", err)
            return
        }

        model.UserCreate(user)

        fmt.Fprintf(w, "OK")
    }
}

func Login(w http.ResponseWriter, r *http.Request){

    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )

    if r.Method == http.MethodPost{

        user := model.User{}

        requestBytes, err := ioutil.ReadAll(r.Body)

        if err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            fmt.Println("io error")
            return
        }

        if err := json.Unmarshal(requestBytes, &user); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            fmt.Println("JSON Unmarshal error:", err)
            return
        }

        user = model.UserLogin(user)

        id := strconv.Itoa(user.UserId)

        resultJson := `{"id":"`+ id + `","name":"` + user.Name + `","token":"` + user.Token + `"}`
        fmt.Println("userlogin" , reflect.TypeOf(resultJson))

        w.WriteHeader(http.StatusOK)
        r.Header.Set("Content-Type", "application/json")

        fmt.Fprintf(w, resultJson)
    }
}
func AutoLogin(w http.ResponseWriter, r *http.Request){

    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )

    if r.Method == http.MethodPost{

        user := model.User{}

        requestBytes, err := ioutil.ReadAll(r.Body)

        if err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            fmt.Println("io error")
            return
        }

        if err := json.Unmarshal(requestBytes, &user); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            fmt.Println("JSON Unmarshal error:", err)
            return
        }

        user = model.UserAutoLogin(user)

        id := strconv.Itoa(user.UserId)

        resultJson := `{"id":"`+ id + `","name":"` + user.Name  + `"}`

        w.WriteHeader(http.StatusOK)
        r.Header.Set("Content-Type", "application/json")

        fmt.Fprintf(w, resultJson)

    }
}