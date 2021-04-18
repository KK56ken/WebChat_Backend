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
func GetFriends(id int)(string,string){
	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	rows,err := db.Query("SELECT receiveFriendId FROM friends WHERE sendFriendId = ?",id)
	defer rows.Close()
	fmt.Println(rows)

	var receiveIds string
	var strid string
	var names string
	cnt := 0

	for rows.Next(){
		receiveId := 0
		if err := rows.Scan(&receiveId); err != nil{
			log.Fatal(err)
		}

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
			if cnt == 0 {
				names = names + name
			}else {
				names = names + "," + name
			}
		}

		// fmt.Println(reflect.TypeOf(receiveId))
		fmt.Println(receiveId)
		fmt.Println(names)

		strid = strconv.Itoa(receiveId)
		// fmt.Println(reflect.TypeOf(strid))
		if cnt == 0 {
			receiveIds = receiveIds + strid
		}else {
			receiveIds = receiveIds + "," + strid
		}
		cnt += 1
	}
	return receiveIds,names
}