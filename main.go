package main

func main() {

	Document = ""
	UserToken = make(map[string]string)
	TokenUser = make(map[string]string)
	Users = make(map[string]string)
	EvSendChan = make(map[string]chan PkgEvent)
	EvRecvChan = make(map[string]chan PkgEvent)

	go editing()
	server()
}
