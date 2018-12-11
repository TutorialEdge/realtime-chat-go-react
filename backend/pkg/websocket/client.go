package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{Type: messageType, Body: string(p)}

		c.Pool.Broadcast <- message

		fmt.Printf("Message Received: %+v\n", message)
		err = c.Conn.WriteJSON(message)
		if err != nil {
			fmt.Println(err)
		}

	}
}
