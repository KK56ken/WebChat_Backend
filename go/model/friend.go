package model

import(
	"database/sql"
	"log"
	"fmt"
	"strconv"
	// "reflect"
	_ "github.com/go-sql-driver/mysql"
)

// type Friend struct {
// 	FriendId int `json:"id"`
// 	Name string `json:"name"`
// 	AvatorURL string `json:"icon"`
// }
func GetFriends(id int)(string){
	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	rows,err := db.Query("SELECT receiveFriendId FROM friends WHERE sendFriendId = ?",id)
	defer rows.Close()
	fmt.Println(rows)

	var strid string
	var json string

	for rows.Next(){

		receiveId := 0

		if err := rows.Scan(&receiveId); err != nil{
			log.Fatal(err)
		}
		strid = strconv.Itoa(receiveId)

		rows2,err := db.Query("SELECT name FROM users WHERE userid = ?",receiveId)
		defer rows2.Close()
		if err != nil{
			log.Fatal(err)
		}
		for rows2.Next(){
			var name string
			if err := rows2.Scan(&name); err != nil{
				log.Fatal(err)
			}
			json += `"{"id":"` + strid + `","name":"` + name + `"},"`

		}
		fmt.Println(json)

		// fmt.Println(reflect.TypeOf(strid))

	}
	return json
}