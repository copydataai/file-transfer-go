package main

import (
	"github.com/leaanthony/clir"
)

func main() {
	cli := clir.NewCli("srvfile", "A server to manage receive and send files", "v0.0.1")
	start := cli.NewSubCommand("start", "Start server to manage")
	start.Action(StartTCPAction)
	stop := cli.NewSubCommand("stop", "stop server")
	stop.Action(StopTCPAction)
	err := cli.Run()
	if err != nil {
		panic(err)
	}
}
