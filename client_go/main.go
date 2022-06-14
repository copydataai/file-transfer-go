package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/leaanthony/clir"
)

// Receive file
func RecvFile(conn net.Conn, channel int) (bool, error) {
	defer conn.Close()
	if _, err := fmt.Fprintf(conn, "Channel: %d GET\n", channel); err != nil {
		fmt.Println("Error sending header channel")
		return false, err
	}
	scan := bufio.NewScanner(conn)
	i := 0
	var filename string
	var fileSize int
	var content bytes.Buffer
	for scan.Scan() {
		line := scan.Text()
		if i == 0 {
			fields := strings.Fields(line)
			filename = fields[1]
		} else if i == 1 {
			fields := strings.Fields(line)
			fileSize, _ = strconv.Atoi(fields[1])
		}
		if line == "" {
			break
		}
		if i != 0 && i != 1 {
			content.Write(scan.Bytes())
		}
		i++
	}
	// TODO verify filesize == content
	fmt.Println(fileSize, content.Len())
	if fileSize == content.Len() {
		// Verify receive all content
		err := ioutil.WriteFile(filename, content.Bytes(), 0644)
		if err != nil {
			fmt.Println("Error Writing: ", err.Error())
			return false, err
		}
	}
	return true, nil
}

// SendFile
func SendFile(conn net.Conn, filename string, channel int) (bool, error) {
	defer conn.Close()
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error file don't exists")
		return false, err
	}
	// REGEX to find filename
	re := regexp.MustCompile(`(\w\s*[^/]*)+\.[a-zA-Z1-9]{1,}$`)
	//  TODO Verify justFilename == ""
	justFilename := re.FindString(filename)

	// Write
	Header := fmt.Sprintf("Channel: %d PUT\nFilename: %s\nFile-Size: %d\n", channel, justFilename, len(file))
	buffer := bytes.NewBufferString(Header)
	if _, err := buffer.Write(file); err != nil {
		fmt.Println("Error Write file in buffer: ", err.Error())
		return false, err

	}
	if _, err := buffer.WriteTo(conn); err != nil {
		fmt.Println("Error to send: ", err.Error())
		return false, err
	}
	return true, nil
}

func main() {
	cli := clir.NewCli("cltfile", "A client to receive and send files", "v0.0.1")

	// Send CMD
	sendCmd := cli.NewSubCommand("send", "Send files to channel")
	s := new(SendCMD)
	sendCmd.IntFlag("channel", "channel to use", &s.channel)
	sendCmd.StringFlag("file", "filename to send", &s.filename)
	sendCmd.Action(s.SendAction)

	// Receive CMD
	r := new(ReceiveCMD)
	receiveCmd := cli.NewSubCommand("receive", "Receive files from a channel")
	receiveCmd.IntFlag("channel", "channel to use", &r.channel)
	receiveCmd.Action(r.ReceiveAction)

	// Catch CLI
	err := cli.Run()
	if err != nil {
		log.Fatal(err)
	}
}
