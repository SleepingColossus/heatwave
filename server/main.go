package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
	upgrader = websocket.Upgrader{} // use default options
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

		err = conn.WriteMessage(mt, strMessage)

		if err != nil {
			log.Println("write:", err)
			break
		}
	}

	fmt.Println("echo end")
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
