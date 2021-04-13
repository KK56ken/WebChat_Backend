package controller

import(
    "net/http"
    "webchat/model"
    // "encoding/json"
    // "io/ioutil"
    "strconv"
)

func Friend(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
	if r.Method == http.MethodGet{
		v := r.URL.Query()

		var id int
		friend := model.Friend{}

		for _, vs := range v{
			id, _ = strconv.Atoi(vs[0])
		}
		friend = model.GetFriends(id)

		
	}

}