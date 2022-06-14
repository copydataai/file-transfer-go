package main

// OBJ to save values from flags
type SendCMD struct {
	channel  int
	filename string
}

func (s *SendCMD) SendAction() error {
	server, err := ConnectTCPServer()
	if err != nil {
		return err
	}
	if ok, err := SendFile(server, s.filename, s.channel); !ok && err != nil {
		return err
	}
	return nil
}
