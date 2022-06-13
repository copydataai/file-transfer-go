package main

type ReceiveCMD struct {
	channel int
}

func (r *ReceiveCMD) ReceiveAction() error {
	client := ConnectTCPServer()
	ok, err := RecvFile(client, r.channel)
	if !ok && err != nil {
		return err
	}
	return nil
}
