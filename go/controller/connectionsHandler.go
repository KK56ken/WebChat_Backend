package controller

import(
    "net/http"
    "webchat/model"
		"fmt"
    "log"
		"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan model.Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Add this lines
	CheckOrigin: func(r *http.Request) bool {
			return true
	},
}

func Connection(w http.ResponseWriter, r *http.Request){

	w.Header().Set( "Access-Control-Allow-Origin", r.RemoteAddr )
	w.Header().Set( "Access-Control-Allow-Credentials", "true" )
	w.Header().Set( "Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization" )
	w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}
	// 関数が終わった際に必ずwebsocketnのコネクションを閉じる
	defer ws.Close()

	// クライアントを新しく登録
	clients[ws] = true

	for {
		var msg model.Message
		// 新しいメッセージをJSONとして読み込みMessageオブジェクトにマッピングする
		err := ws.ReadJSON(&msg)
		if err != nil {
				log.Printf("error: %v", err)
				delete(clients, ws)
				break
		}
		fmt.Println(msg)
		// 新しく受信されたメッセージをブロードキャストチャネルに送る
		broadcast <- msg
	}
}
func HandleMessages() {
	for {
			// ブロードキャストチャネルから次のメッセージを受け取る
			msg := <-broadcast
			// 現在接続しているクライアント全てにメッセージを送信する
			for client := range clients {
					err := client.WriteJSON(msg)
					if err != nil {
							log.Printf("error: %v", err)
							client.Close()
							delete(clients, client)
					}
			}
	}
}