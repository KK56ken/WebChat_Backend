package model

import(
	"database/sql"
	"log"
	"fmt"
	// "strconv"
	// "reflect"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

type Chat struct {
	Id int `json:"Id"`
	SendUserId int `json:"sendUserId"`
	ReceiveUserId int `json:"receiveUserId"`
	Message string `json:"message"`
	Time string `json:"time"`
}

func GetChat(userId int, ids []string)([]string){

	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")

	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	inIds := "IN ('" + strings.Join(ids, "', '") + "')";

	sql := "SELECT * FROM messages WHERE sendUserId = ? AND receiveUserId " + inIds + " UNION SELECT * FROM messages WHERE receiveUserId = ? AND sendUserId " + inIds + " ORDER BY time ASC"

	message := new(Chat)
	rows,err := db.Query(sql,userId,userId)
	defer rows.Close()

	var json []string

	for rows.Next(){

		if err := rows.Scan(&message.Id,&message.SendUserId,&message.ReceiveUserId,&message.Message,&message.Time); err != nil{
			log.Fatal(err)
		}
		sname := GetFriendName(message.SendUserId)
		rname := GetFriendName(message.ReceiveUserId)

		json = append(json,`{"avater":"''` + `","sendUserName":"`+ sname + `","receiveUserName":"` + rname + `","message":"` + message.Message + `","time":"` + message.Time + `"}`)

		fmt.Println(message.SendUserId,message.ReceiveUserId,message.Message,message.Time)
		fmt.Println(json)
	}

	return json
}