package main

import (
	"fmt"
	"net"
)

// Connet TCP server
func ConnectTCPServer() (net.Conn, error) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error to connect")
		return nil, err
	}
	return conn, nil
}
