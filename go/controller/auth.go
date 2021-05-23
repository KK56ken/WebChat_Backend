package controller

import(
	"net/http"
	// "webchat/model"
	// "fmt"
	// "log"
	"encoding/json"
	"os"
	"github.com/auth0/go-jwt-middleware"
  "github.com/form3tech-oss/jwt-go"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

func Private(w http.ResponseWriter, r *http.Request){

	w.Header().Set( "Access-Control-Allow-Origin", r.RemoteAddr )
	w.Header().Set( "Access-Control-Allow-Credentials", "true" )
	w.Header().Set( "Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization" )
	w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )

	post := &post{
		Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
		Tag:   "Go",
		URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
	}
	json.NewEncoder(w).Encode(post)


}
func Public(w http.ResponseWriter, r *http.Request){

	w.Header().Set( "Access-Control-Allow-Origin", r.RemoteAddr )
	w.Header().Set( "Access-Control-Allow-Credentials", "true" )
	w.Header().Set( "Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization" )
	w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )

	post := &post{
		Title: "VueCLIからVue.js入門①【VueCLIで出てくるファイルを概要図で理解】",
		Tag:   "Vue.js",
		URL:   "https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
	}
	json.NewEncoder(w).Encode(post)


}


// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})