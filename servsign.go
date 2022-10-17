package main

import (
	"fmt"
	"net"
)

func servSign(conn net.Conn, bs []byte) {
	// Obtener usuario
	user, _ := BytesToUser(bs)

	// Confirmar disponibilidad del username
	MutexUsers.StartReader()
	_, ok := Users[user.name]
	MutexUsers.EndReader()

	if !ok {
		// Registrar usuario
		MutexUsers.StartWrite()
		Users[user.name] = user.pass
		MutexUsers.EndWrite()

		// Registrar Canales de comunicacion
		MutexChann.StartWrite()
		EvRecvChan[user.name] = make(chan PkgEvent, 32)
		EvSendChan[user.name] = make(chan PkgEvent, 32)
		MutexChann.EndWrite()

		fmt.Printf("[OK] Usuario '%s' registrado!\n", user.name)
		conn.Write([]byte{1})
	} else {
		fmt.Printf("[ERROR] Usuario '%s' Ya existe!\n", user.name)
		conn.Write([]byte{0})
	}
	conn.Close()
}
