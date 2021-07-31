package main

import (
	"encoding/json"
	"flag"
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

	sendCh      = make(chan *ChannelMessage)
	broadcastCh = make(chan *ChannelMessage)
)

// server constants
const (
	// binary message type
	binMt = 2
)

// game state vars
var (
	gameState = game.NewGameState()
)

// game state constants
const (
	// 30 fps
	tickRate time.Duration = time.Second / 30
)

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, strMessage, err := conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}

		var message ClientMessage
		if err = json.Unmarshal(strMessage, &message); err != nil {
			log.Println("unmarshal:", err)
			break
		}

		log.Printf("event type: %d\t event body: %s", message.MessageType, message.MessageBody)

		var response ServerMessage
		switch message.MessageType {
		case JoinGame:
			clientId := uuid.New().String()

			// add new active player
			clients[clientId] = conn

			// create game state snapshot
			gameStateUpdate := gameState.AddPlayer(clientId)
			gameStateUpdate.SetClientId(clientId)

			// send snapshot to current connecting player
			response = newServerMessage(InitNewPlayer, gameStateUpdate)
			sendCh <- newChannelMessage(response, conn)

			// TODO discard?
			// change message type before broadcasting to all other players
			gameStateUpdate.ClearClientId()
			response.MessageType = GameStateUpdated
			broadcastCh <- newChannelMessage(response, nil)

		case LeaveGame:
			clientId, ok := message.MessageBody["clientId"]

			if !ok {
				log.Println("client id is blank")
				return
			}

			delete(clients, clientId)
			gameState.MarkPlayerForDeletion(clientId)

		case Move:
			playerId := message.MessageBody["clientId"]
			directionX, _ := strconv.Atoi(message.MessageBody["x"])
			directionY, _ := strconv.Atoi(message.MessageBody["y"])

			err := gameState.PlayerMove(playerId, directionX, directionY)

			if err != nil {
				log.Println(err)
			}

		case Shoot:
			actorId := message.MessageBody["clientId"]
			err := gameState.PlayerShoot(actorId)

			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

// TODO move logic to gamestate.go
func startMoveActorsTask() {
	lastTime := time.Now()
	for {
		if gameState.IsActive() {
			t := time.Now()
			dt := t.Sub(lastTime)

			// is it time for the next frame?
			if dt >= tickRate {
				// create payload
				update := gameState.Update()
				response := newServerMessage(GameStateUpdated, update)

				// broadcast event
				broadcastCh <- newChannelMessage(response, nil)

				// set timer for next iteration
				lastTime = t
			}
		}
	}
}

func startSendListener() {
	for {
		msg := <-sendCh

		strMessage, err := json.Marshal(msg.Message)

		if err != nil {
			log.Println("marshall error:", err)
			return
		}

		err = msg.Connection.WriteMessage(binMt, strMessage)

		if err != nil {
			log.Println("write:", err)
		}
	}
}

func startBroadcastListener() {
	for {
		msg := <-broadcastCh
		for id, conn := range clients {
			msg.Message.MessageBody.SetClientId(id)
			msg.Connection = conn
			sendCh <- msg
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", handleWS)

	go startSendListener()
	go startBroadcastListener()
	go startMoveActorsTask()

	log.Println("listening on:", *addr)
	log.Fatal(http.ListenAndServe(*addr, logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
