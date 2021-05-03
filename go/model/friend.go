package model

import(
	"database/sql"
	"log"
	"fmt"
	"strconv"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
)

type Friend struct {
	FriendId int `json:"id"`
	Name string `json:"name"`
	AvatorURL string `json:"icon"`
}
func GetFriends(id int)([]string){
	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	rows,err := db.Query("SELECT receiveFriendId FROM friends WHERE sendFriendId = ?",id)
	defer rows.Close()
	fmt.Println(rows)

	var json []string

	for rows.Next(){
		friend := new(Friend)

		if err := rows.Scan(&friend.FriendId); err != nil{
			log.Fatal(err)
		}
		strid := strconv.Itoa(friend.FriendId)

		rows2,err := db.Query("SELECT name FROM users WHERE userid = ?", friend.FriendId)
		defer rows2.Close()
		if err != nil{
			log.Fatal(err)
		}
		for rows2.Next(){
			if err := rows2.Scan(&friend.Name); err != nil{
				log.Fatal(err)
			}
			json = append(json,`{"id":"`+ strid + `","name":"` + friend.Name + `"}`)
		}
		fmt.Println(json)
		fmt.Println(reflect.TypeOf(json))
	}
	return json
}