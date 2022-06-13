package main

import (
	"bytes"
	"fmt"
	"net"
)

type Channel struct {
	Number          int
	IncomingDevices chan net.Conn
	LeavingDevices  chan net.Conn
	sendFiles       chan bytes.Buffer
}

var (
	Channel0 = Channel{0, make(chan net.Conn, 10), make(chan net.Conn, 10), make(chan bytes.Buffer)}
	Channel1 = Channel{1, make(chan net.Conn, 10), make(chan net.Conn, 10), make(chan bytes.Buffer)}
	Channel2 = Channel{2, make(chan net.Conn, 10), make(chan net.Conn, 10), make(chan bytes.Buffer)}
)

// SendFile require conn and content
// conn is a connection Waiting in IncomingDevices
func SendFile(conn net.Conn, content bytes.Buffer, channel int) {
	length, err := conn.Write(content.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println("Length send", length)
	switch channel {
	case 0:
		Channel0.LeavingDevices <- conn
	case 1:
		Channel1.LeavingDevices <- conn
	case 2:
		Channel2.LeavingDevices <- conn
	}
	conn.Close()
}

// Manager from channels
func Broadcast(channel *Channel) {
	// Use Map for no to take other process to delete in a Slice
	clients := make(map[net.Conn]int)
	for {
		select {
		case file := <-channel.sendFiles:
			for client := range clients {
				go SendFile(client, file, channel.Number)
			}
		case newClient := <-channel.IncomingDevices:
			clients[newClient] = channel.Number
		case leavingClient := <-channel.LeavingDevices:
			delete(clients, leavingClient)
		}
	}
}

func main() {
	StartTCPServer()
}
