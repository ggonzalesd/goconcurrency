package main

import (
	"net"
)

func servRecv(conn net.Conn, bs []byte) {
	token := string(bs[:32])
	MutexToken.StartReader()
	user, ok := TokenUser[token]
	MutexToken.EndReader()
	if ok {
		conn.Write([]byte{1})
		buffer := make([]byte, 1024)
		MutexChann.StartReader()
		c := EvRecvChan[user]
		MutexChann.EndReader()
		for {
			_, err := conn.Read(buffer)
			if err != nil {
				break
			}
			pkge, _ := BytesToPkgEvent(buffer)
			c <- pkge
			conn.Write([]byte{1})
		}
	} else {
		conn.Write([]byte{0})
	}
}
