package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func HandleConn(client net.Conn) {
	i := 0
	scan := bufio.NewScanner(client)
	var channel int
	var isPUT bool
	content := bytes.Buffer{}
	for scan.Scan() {
		line := scan.Text()
		// Type Send o Receive
		if i == 0 {
			fields := strings.Fields(line)
			channel, _ = strconv.Atoi(fields[1])
			if fields[2] == "PUT" {
				isPUT = true
			} else {
				break
			}
		}

		if line == "" {
			break
		}
		if i != 0 && i > 2 {
			content.Write(scan.Bytes())
			content.Write([]byte("\n"))
		}
		i++
	}
	if isPUT {
		switch channel {
		case 0:
			Channel0.sendFiles <- content
		case 1:
			Channel1.sendFiles <- content
		case 2:
			Channel2.sendFiles <- content
		}
		client.Close()
		return
	}
	switch channel {
	case 0:
		Channel0.IncomingDevices <- client
	case 1:
		Channel1.IncomingDevices <- client
	case 2:
		Channel2.IncomingDevices <- client
	}
}

func StartTCPServer() {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	go Broadcast(&Channel0)
	go Broadcast(&Channel1)
	go Broadcast(&Channel2)
	for {
		fmt.Println("Wait a client")
		client, err := conn.Accept()
		if err != nil {
			log.Println("error Acceting", err.Error())
			continue
		}
		go HandleConn(client)
	}
}
