package controller

import(
    "net/http"
    "webchat/model"
		"fmt"
    "strconv"
		// "strings"
)

func Friend(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
	if r.Method == http.MethodGet{
		v := r.URL.Query()

		var id int
		var friendsJson string

		for _, vs := range v{
			id, _ = strconv.Atoi(vs[0])
		}
		friendsJson = model.GetFriends(id)

		fmt.Fprintf(w, friendsJson)
	}

}