package model

import(
	"database/sql"
	"log"
	"fmt"
	// "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

type Friend struct {
	FriendId int `json:"id"`
	Name string `json:"name"`
	AvatorURL string `json:"icon"`
}
func GetFriends(id int){
	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	rows,err := db.Query("SELECT receiveFriendId FROM friends WHERE sendFriendId = ?",id)

	defer rows.Close()

	// var friendResult []Friend
	// friend := Friend{}
	var receiveIds []int

	for rows.Next(){
		receiveId := 0
		if err := rows.Scan(&receiveId); err != nil{
			log.Fatal(err)
		}
		receiveIds = append(receiveIds, receiveId)
	}
	for i := 0; i < len(receiveIds); i++ {
		fmt.Println("receiveFriendId: ", receiveIds[i])
	}



}