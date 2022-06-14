package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func ReadInformation() error {
	tmpDirs, err := ioutil.ReadDir("/tmp/")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("^file-go-cache.+")
	var tmpFile string
	for _, value := range tmpDirs {
		if length := re.FindAllString(value.Name(), 1); len(length) > 0 {
			tmpFile = length[0]
		}
	}
	pid, err2 := os.ReadFile("/tmp/" + tmpFile)
	if err2 != nil {
		panic(err2)
	}
	pidInt, _ := strconv.Atoi(string(pid))
	pidNew := os.Process{
		Pid: pidInt,
	}
	pidNew.Kill()
	if err := os.Remove("/tmp/" + tmpFile); err != nil {
		return err
	}
	return nil
}

func StopTCPAction() error {
	err := ReadInformation()
	if err != nil {
		fmt.Println("Don't exists a server to stop")
		return err
	}
	fmt.Println("Server stop(i kill the server)")
	return nil
}
