package main

type ReceiveCMD struct {
	channel int
}

func (r *ReceiveCMD) ReceiveAction() error {
	client, err := ConnectTCPServer()
	if err != nil {
		return err
	}
	ok, err1 := RecvFile(client, r.channel)
	if !ok && err1 != nil {
		return err1
	}
	return nil
}
