package model

import(
	"database/sql"
	"log"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserId int `json:"user"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
	Token string `json:"token"`
	AvatorURL string `json:"avatorurl"`
}

func UserCreate(user User){
	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	token, err := TokenCreate(user)

	user.Token = token

	_, err = db.Exec("INSERT INTO users(email, name, password, token) VALUES(?,?,?,?)",user.Email, user.Name, user.Password, user.Token)
	if err != nil{
		log.Fatal(err)
	}

}
func TokenCreate(user User)(string, error){
	var err error

	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss": "__init__",//JWTの発行者が入る（文字列（__init__）は任意）
	})

	// spew.Dump(token)

	tokenString, err := token.SignedString([]byte(secret))

	fmt.Println("-----------------------------")
  fmt.Println("tokenString:", tokenString)

	if err != nil{
		log.Fatal(err)
	}
	return tokenString, nil
}

func UserLogin(user User)(User){
	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	rows,err := db.Query("SELECT userid, name, token  FROM users WHERE Email = ? AND Password = ?",user.Email, user.Password)

	defer rows.Close()

	for rows.Next(){
		if err := rows.Scan(&user.UserId,&user.Name,&user.Token); err != nil{
			log.Fatal(err)
		}
	}
	return user
}

func GetFriendName(id int) string{
	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT name FROM users WHERE userid = ?"

	rows,err := db.Query(sql, id)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println(rows)
	var friendName string
	for rows.Next(){
		if err := rows.Scan(&friendName); err != nil{
			log.Fatal(err)
		}
		fmt.Println(friendName)
	}
	return friendName
}
func UserAutoLogin(user User)(User){
	db,err := sql.Open("mysql","root:root@tcp(mysql_host:3306)/chatDB")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT userid, email, name FROM users WHERE token = ?"
	rows,err := db.Query(sql, user.Token)
	defer rows.Close()

	for rows.Next(){
		if err := rows.Scan(&user.UserId, &user.Email, &user.Name); err != nil{
			log.Fatal(err)
		}
	}
	return user
}