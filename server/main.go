package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

	sendCh      = make(chan *ServerMessage)
	broadcastCh = make(chan *ServerMessage)
)

// game state vars
var (
	players = make(map[string]*game.Actor, 0)
	enemies = make(map[string]*game.Actor, 0)
)

// game state constants
const (
	screenWidth  = 1920
	screenHeight = 1080

	// 30 fps
	tickRate time.Duration = 1000 / 30
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

		var message ClientMessage
		if err = json.Unmarshal(strMessage, &message); err != nil {
			log.Println("unmarshal:", err)
			break
		}

		log.Printf("event type: %d", message.MessageType)
		log.Printf("event body: %s", message.MessageBody)

		var response *ServerMessage
		switch message.MessageType {
		case JoinGame:
			clientId := uuid.New().String()

			// add new active player
			clients[clientId] = conn
			players[clientId] = game.NewActor(clientId, game.DefaultPlayerPosition())

			// create reply with coords of player
			msgBody := []map[string]string{{
				"clientId":  clientId,
				"actorType": strconv.Itoa(game.Player),
				"x":         "100",
				"y":         "100",
			}}

			// send join response to currently connecting player
			response = newServerMessage(SelfConnected, msgBody)
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
			delete(players, clientId)

			msgBody := []map[string]string{{
				"clientId": clientId,
			}}

			response = newServerMessage(PlayerDisconnected, msgBody)
			fmt.Printf("player has left: %s\n", clientId)

		case Move:
			updatePlayerDirection(message.MessageBody)
		}

		if response != nil {
			go broadcast(mt, *response)
		}
	}

	fmt.Println("echo end")
}

func updatePlayerDirection(dir map[string]string) {
	actorId := dir["clientId"]
	directionX, _ := strconv.Atoi(dir["x"])
	directionY, _ := strconv.Atoi(dir["y"])

	direction := game.NewVector2(directionX, directionY)

	actor, ok := players[actorId]

	if !ok {
		fmt.Println("unknown actor", actorId)
		return
	}

	actor.Direction = direction
}

func moveActors() {
	lastTime := time.Now()
	for {
		t := time.Now()
		dt := t.Sub(lastTime)

		if len(players) > 0 {
			// is it time for the next frame?
			if dt >= tickRate {

				// container for new positions
				var messageBody []map[string]string

				// move all players
				for _, actor := range players {
					actor.Move()
					messageBody = append(messageBody, actor.ToMap())
				}

				// move all enemies
				for _, enemy := range enemies {
					enemy.Move()
					messageBody = append(messageBody, enemy.ToMap())
				}

				// create payload
				response := newServerMessage(ActorsMoved, messageBody)

				// broadcast event
				go broadcast(0, *response)

				// set timer for next iteration
				lastTime = t
			}
		}
	}
}

func send(mt int, message ServerMessage, conn *websocket.Conn) {
	strMessage, err := json.Marshal(message)

	if err != nil {
		log.Println("marshall error:", err)
		return
	}

	err = conn.WriteMessage(mt, strMessage)

	if err != nil {
		log.Println("write:", err)
	}
}

func broadcast(mt int, message ServerMessage) {
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

	//go moveActors()

	fmt.Println("listening on:", *addr)
	log.Fatal(http.ListenAndServe(*addr, logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
