package controller

import(
    "net/http"
    "webchat/model"
		"fmt"
    "strconv"
		// "strings"
)

// func (is friendsInt) toString() []string {
//   f := make([]string, len(is))
//   for i, v := range is {
//     f[i] = string(v)
//   }
//   return f
// }

func Friend(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
	if r.Method == http.MethodGet{
		v := r.URL.Query()

		var id int

		var ids string
		var names string

		for _, vs := range v{
			id, _ = strconv.Atoi(vs[0])
		}
		ids,names = model.GetFriends(id)

		// fmt.Println(friends)
		friendsJson := `{"ids":"` + ids + `","names":"` + names + `"}`

		fmt.Fprintf(w, friendsJson)
	}

}