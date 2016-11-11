package main

import (
	"fmt"
	"log"
	"io"
	"flag"
	"time"

	"golang.org/x/net/websocket"
	"./base"
)

var origin = "http://localhost/"
var url = "ws://localhost:8080/echo"

func main() {
	code := flag.String("id", "", "")
	flag.Parse()

	var err error
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(time.Millisecond * 500)
	var i int
	for t := range ticker.C {
		log.Println(t)
		i = i + 1
		message := base.Message{Author: *code, Body: fmt.Sprintf("message %d", i)}
		websocket.JSON.Send(ws, message)
		fmt.Printf("Send: %s\n", message)

		var msg base.Message
		err = websocket.JSON.Receive(ws, &msg)
		if err == io.EOF {
			fmt.Println("EOF")
			//c.doneCh <- true
		} else if err != nil {
			fmt.Println(err)
			//c.server.Err(err)
		} else {
			log.Println("received", msg)
		}

		fmt.Printf("Receive: %s\n", msg)
	}
}
