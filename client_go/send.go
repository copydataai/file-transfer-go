package main

import (
	"fmt"
)

type SendCMD struct {
	channel  int
	filename string
}

func (s *SendCMD) SendAction() error {
	server := ConnectTCPServer()
	ok := SendFile(server, s.filename, s.channel)
	if !ok {
		return fmt.Errorf("Error to send file")
	}
	return nil
}
