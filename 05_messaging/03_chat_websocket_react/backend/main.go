package main

import (
	"fmt"
	"net/http"

	"github.com/miguelsilva5989/chat-websocket/pkg/websocket"
)

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	websocket.Reader(ws)
}

func setupRoutes() {
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
