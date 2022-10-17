package main

import (
	"fmt"
	"net"
	"os"
)

func server() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error Listening: ", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)

	for {
		connection, err := server.Accept()
		if err == nil {
			processClient(connection)
		}
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	_, err := connection.Read(buffer)
	if err == nil {
		switch buffer[0] {
		case 1:
			go servSign(connection, buffer[1:])
		case 2:
			go servLog(connection, buffer[1:])
		case 3:
			go servRecv(connection, buffer[1:])
		case 4:
			go servSend(connection, buffer[1:])
		default:
			connection.Close()
		}
	}
}
