package controller

import(
    "net/http"
    "webchat/model"
		"fmt"
    "strconv"
		"log"
		"encoding/json"
)

func Friend(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
	if r.Method == http.MethodGet{
		v := r.URL.Query()

		var id int

		for _, vs := range v{
			id, _ = strconv.Atoi(vs[0])
		}
		friends := model.GetFriends(id)

		jsonBytes, err := json.Marshal(friends)

		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			log.Fatal(err)
			return
		}

		jsonString := string(jsonBytes)

		w.WriteHeader(http.StatusOK)
    r.Header.Set("Content-Type", "application/json")

		fmt.Fprintf(w, jsonString)
	}

}