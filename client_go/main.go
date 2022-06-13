package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/leaanthony/clir"
)

// Receive file
func RecvFile(conn net.Conn, channel int) (bool, error) {
	defer conn.Close()
	length, err := fmt.Fprintf(conn, "Channel: %d GET\n", channel)
	if err != nil {
		fmt.Println("Error sending header channel")
		return false, err
	}
	fmt.Println("Length:", length)
	scan := bufio.NewScanner(conn)
	i := 0
	var filename string
	var fileSize int
	var content bytes.Buffer
	for scan.Scan() {
		if i == 0 {
			fields := strings.Fields(scan.Text())
			filename = fields[1]
		} else if i == 1 {
			fields := strings.Fields(scan.Text())
			fileSize, _ = strconv.Atoi(fields[1])
		}
		if scan.Text() == "" {
			break
		}
		if i != 0 && i != 1 {
			content.Write(scan.Bytes())
		}
		i++
	}
	if fileSize > content.Len() {
		err := ioutil.WriteFile(filename, content.Bytes(), 0644)
		if err != nil {
			fmt.Println("My file error: ", err.Error())
			return false, err
		}
	}
	return true, nil
}

// SendFile
func SendFile(conn net.Conn, filename string, channel int) bool {
	defer conn.Close()
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error file don't exists")
		return false
	}
	re := regexp.MustCompile(`(\w\s*[^/]*)+\.[a-zA-Z1-9]{1,}$`)
	justFilename := re.FindString(filename)
	Header := fmt.Sprintf("Channel: %d PUT\nFilename: %s\nFile-Size: %d\n", channel, justFilename, len(file))
	buffer := bytes.NewBufferString(Header)
	if err != nil {
		fmt.Println("Error send Header: ", err.Error())
		return false
	}
	writting, err := buffer.Write(file)
	if err != nil {
		fmt.Println("Error Write file in buffer: ", err.Error())
		return false

	}
	fmt.Println("Writting:", writting)
	length, err := buffer.WriteTo(conn)
	if err != nil {
		fmt.Println("Error Write conn: ", err.Error())
		return false
	}
	fmt.Println("Length from buffer send: ", length)
	return true
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
	err := cli.Run()
	if err != nil {
		panic(err)
	}
}
