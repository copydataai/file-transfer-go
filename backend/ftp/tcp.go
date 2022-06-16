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

// A handle Conn from get or put and channel
func HandleConn(client net.Conn) {
	i := 0
	scan := bufio.NewScanner(client)
	var channel int
	var isPUT bool
	content := bytes.Buffer{}
	for scan.Scan() {
		line := scan.Text()
		// Type Send o Receive
		// And Channel
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
			// Wait Forever
			break
		}
		if i != 0 {
			content.Write(scan.Bytes())
			if i < 3 {
				// Just for separe filename & filesize
				// and content is continuos
				// to don't alterate the file
				// Make this 01010101\n01010101
				content.Write([]byte("\n"))
			}
		}
		i++
	}
	if isPUT {
		// I leave this by don't think other reference(and search)
		// look how JS
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
