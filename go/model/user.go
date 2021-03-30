package model

import(
	"database/sql"
	"log"
)

type User struct {
	UserId int `json:"user"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	Token string `json:"token"`
	AvatorURL string `json:"avatorurl"`
}


func UserCreate(email string, name string, password string){
	db,_ := sql.Open("mysql","user=root password=root dbname=chatDB")
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO users(email, name, password, token) VALUES(?,?,?,?)")
	if err != nil{
		log.Fatal(err)
	}
	token, err := TokenCreate()

	ins.Exec(email, name, password, token)
}
func TokenCreate(){



}