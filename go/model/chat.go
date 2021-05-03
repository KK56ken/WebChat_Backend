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
	SendUserId string `json:"sendUserId"`
	ReceiveUserId string `json:"receiveUserId"`
	Message string `json:"message"`
	Time string `json:"time"`
}

func GetChat(userId int, ids []string)(string){

	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")

	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	inIds := "IN ('" + strings.Join(ids, "', '") + "')";

	sql := "SELECT * FROM messages WHERE sendUserId = ? AND receiveUserId " + inIds + " UNION SELECT * FROM messages WHERE receiveUserId = ? AND sendUserId " + inIds + " ORDER BY time ASC"
	// sql := "SELECT * FROM messages WHERE sendFriendId = ?"
	// fmt.Println("debug1",sql)
	//"SELECT sendUserId, receiveUserId, message, time FROM messages WHERE sendUserId = ?"

	message := new(Chat)
	rows,err := db.Query(sql,userId,userId)
	defer rows.Close()

	var messages = make([][]string,len(ids))

	for rows.Next(){

		if err := rows.Scan(&message.Id,&message.SendUserId,&message.ReceiveUserId,&message.Message,&message.Time); err != nil{
			log.Fatal(err)
		}
		fmt.Println(ids)

		for i, v := range ids{
			fmt.Println(i)
			if v == message.SendUserId || v == message.ReceiveUserId{
				messages[i] = append(messages[i],"{avater:'',receiveUserName:'',sendUserName'',message:"+ message.Message + "},")
			}
			fmt.Println(i)
		}

		fmt.Println(message.SendUserId,message.ReceiveUserId,message.Message,message.Time)
	}

	fmt.Println(messages)
	fmt.Println(messages[0])
	fmt.Println(messages[1])

	var responseMessage string
	for _, v := range ids{
		responseMessage += "id" + v + ":["
		for _,v2 := range messages{
			for _,v3 := range v2{
				responseMessage += v3
			}
		}
		responseMessage += "],"
	}

	return responseMessage
}