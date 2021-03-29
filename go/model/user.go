package main

import(
	"database/sql"
)

var DbConnection *sql.DB

func main(){
	DbConnection,_ := sql.Open("mysql",)
	defer DbConnection.Close()
}