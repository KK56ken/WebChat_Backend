package model

// import(
// 	"database/sql"
// )

type Message struct {
	UserId int `json:"user"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	Token string `json:"token"`
	Avatorurl string `json:"avatorurl"`
}