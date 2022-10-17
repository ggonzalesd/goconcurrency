package main

import "net"

func servSend(conn net.Conn, bs []byte) {
	token := string(bs[:32])
	MutexToken.StartReader()
	user, ok := TokenUser[token]
	MutexToken.EndReader()
	if ok {
		conn.Write([]byte{1})
		buffer := make([]byte, 1024)
		MutexChann.StartReader()
		c := EvSendChan[user]
		MutexChann.EndReader()
		for {
			pkge := <-c
			conn.Write(PkgEventToBytes(pkge))
			_, err := conn.Read(buffer[:1])
			if err != nil {
				break
			}
		}
	} else {
		conn.Write([]byte{0})
	}
}
