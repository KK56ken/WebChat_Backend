package controller

import(
    "fmt"
    "net/http"
    "webchat/model"
    "strconv"
    // "encoding/json"
    // "log"
)

func Chat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )

    if r.Method == http.MethodGet{

        v := r.URL.Query()

		var id int
        //ログインユーザーのID取得
		for _, vs := range v{
			id, _ = strconv.Atoi(vs[0])
        }

        //データベースでidと一致している
        ids := model.GetFriendsId(id)

        messages := model.GetChat(id, ids)

        // jsonBytes, err := json.Marshal(messages)

		// if err != nil {
		// 	w.WriteHeader(http.StatusServiceUnavailable)
		// 	log.Fatal(err)
		// 	return
		// }

		// jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)
        r.Header.Set("Content-Type", "application/json")

        fmt.Fprintf(w, messages)
    }else if r.Method == http.MethodPost{

        // user := model.User{}

        // requestBytes, err := ioutil.ReadAll(r.Body)

        // if err != nil {
        //     w.WriteHeader(http.StatusServiceUnavailable)
        //     fmt.Println("io error")
        //     return
        // }

        // if err := json.Unmarshal(requestBytes, &user); err != nil {
        //     w.WriteHeader(http.StatusServiceUnavailable)
        //     fmt.Println("JSON Unmarshal error:", err)
        //     return
        // }

        // user = model.UserLogin(user)

        // id := strconv.Itoa(user.UserId)

        // resultJson := `{"id":"`+ id + `","name":"` + user.Name + `","token":"` + user.Token + `"}`
        // fmt.Println("userlogin" , reflect.TypeOf(resultJson))

        // w.WriteHeader(http.StatusOK)
        // r.Header.Set("Content-Type", "application/json")

        // fmt.Fprintf(w, resultJson)
    }
}