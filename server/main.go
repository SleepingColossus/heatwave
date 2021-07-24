package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SleepingColossus/heatwave/game"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// server vars
var (
	addr     = flag.String("addr", ":8080", "http service address")
	upgrader = websocket.Upgrader{} // use default options

	// active players
	clients = make(map[string]*websocket.Conn)
)

// game state vars
var (
	actors = make([]game.Actor, 0)
	broadcastCh = make(chan *Message)
)

// game state constants
const (
	screenWidth  = 1920
	screenHeight = 1080
)

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("echo start")

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		mt, strMessage, err := conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}

		var message Message
		if err = json.Unmarshal(strMessage, &message); err != nil {
			log.Println("unmarshal:", err)
			break
		}

		log.Printf("event type: %d", message.MessageType)
		log.Printf("event body: %s", message.MessageBody)

		var response *Message
		switch message.MessageType {
		case JoinGame:
			clientId := uuid.New()

			// add new active player
			clients[clientId.String()] = conn

			// create reply with coords of player
			msgBody := map[string]string{
				"clientId":  clientId.String(),
				"actorType": strconv.Itoa(game.Player),
				"x":         "1000",
				"y":         "500",
			}

			// send join response to currently connecting player
			response = newMessage(SelfConnected, msgBody)
			go send(mt, *response, conn)

			// change message type before broadcasting to all other players
			response.MessageType = PlayerConnected
			fmt.Printf("player has joined: %s\n", clientId)

		case LeaveGame:
			clientId, ok := message.MessageBody["clientId"]

			if !ok {
				fmt.Println("client id is blank")
				return
			}

			delete(clients, clientId)

			msgBody := map[string]string {
				"clientId": clientId,
			}

			response = newMessage(PlayerDisconnected, msgBody)
			fmt.Printf("player has left: %s\n", clientId)
		}

		go broadcast(mt, *response)
	}

	fmt.Println("echo end")
}

func send(mt int, message Message, conn *websocket.Conn) {
	strMessage, err :=json.Marshal(message)

	if err != nil {
		log.Println("marshall error:", err)
		return
	}

	err = conn.WriteMessage(mt, strMessage)

	if err != nil {
		log.Println("write:", err)
	}
}

func broadcast(mt int, message Message) {
	fmt.Println("broadcast start")

	for _, conn := range clients {
		strMessage, err := json.Marshal(message)

		if err != nil {
			log.Println("marshall error:", err)
			return
		}

		err = conn.WriteMessage(mt, strMessage)

		if err != nil {
			log.Println("write:", err)
			break
		}
	}

	fmt.Println("broadcast end")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", echo)

	fmt.Println("listening on:", *addr)
	log.Fatal(http.ListenAndServe(*addr, logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
