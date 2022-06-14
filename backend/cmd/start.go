package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func SaveInformation(pid *int) error {
	tmpFile, err := ioutil.TempFile("", "file-go-cache")
	if err != nil {
		return err
	}
	if _, err := tmpFile.WriteString(fmt.Sprintf("%d", *pid)); err != nil {
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}
	return nil
}

func StartTCPAction() error {
	path := "/home/copy/Documents/truora/file_transfer/backend/ftp/ftp"
	cmd := exec.Cmd{Path: path}
	err := cmd.Start()
	if err != nil {
		return err
	}
	if err := SaveInformation(&cmd.Process.Pid); err != nil {
		return err
	}

	return nil
}
