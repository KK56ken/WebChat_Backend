package controller

import(
    "fmt"
    "net/http"
    "webchat/model"
    "strconv"
    "encoding/json"
    "log"
		"io/ioutil"
)

func Message(w http.ResponseWriter, r *http.Request) {
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
		// fmt.Println(ids)

		messages := model.GetMessages(id, ids)

		jsonBytes, err := json.Marshal(messages)

		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")

		fmt.Fprintf(w, jsonString)

	}else if r.Method == http.MethodPost{

		message := model.Message{}

		requestBytes, err := ioutil.ReadAll(r.Body)

		if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Println("io error")
				return
		}

		if err := json.Unmarshal(requestBytes, &message); err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				fmt.Println("JSON Unmarshal error:", err)
				return
		}
		// fmt.Println(message)

		model.PostMessage(message)

		sname := model.GetFriendName(message.SendUserId)
		rname := model.GetFriendName(message.ReceiveUserId)

		resultJson := `{"avater":"''` + `","sendUserName":"`+ sname + `","receiveUserName":"` + rname + `","message":"` + message.Message + `","time":"` + message.Time + `"}`
		// fmt.Println(message.SendUserId,message.ReceiveUserId,message.Message,message.Time)

		w.WriteHeader(http.StatusOK)
		r.Header.Set("Content-Type", "application/json")

		fmt.Fprintf(w, resultJson)
	}
}