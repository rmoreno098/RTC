package main

import (
	"github.com/gorilla/websocket"
)

// Message struct
type Message struct {
	Type    string `json:"type"`
	ID      string `json:"id"`
	Message string `json:"message"`
}

type MapResponse struct {
	ID string `json:"id"`
}

// Client struct
type Client struct {
	conn *websocket.Conn
	id   string
	send chan []byte
}
