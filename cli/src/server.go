package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
        "./base"
)

func echoHandler(ws *websocket.Conn) {
	for {
		var msg base.Message
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Receive", msg)

		websocket.JSON.Send(ws, msg)
		fmt.Printf("Send: %s\n", msg)
	}
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
