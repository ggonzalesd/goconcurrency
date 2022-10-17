package main

import (
	"net"
	"strings"
)

func servLog(conn net.Conn, bs []byte) {
	// Obtener usuario
	user, _ := BytesToUser(bs)

	// Confirmar existencia del usuario
	MutexUsers.StartReader()
	pass, ok := Users[user.name]
	MutexUsers.EndReader()
	if ok {
		ok = strings.Compare(user.pass, pass) == 0
	}

	if ok {
		// Generate Token
		MutexToken.StartWrite()
		token, ok := UserToken[user.name]
		if ok {
			delete(TokenUser, token)
		}
		token = generateToken()
		UserToken[user.name] = token
		TokenUser[token] = user.name
		MutexToken.EndWrite()

		MutexDoc.StartReader()
		doc := string(Document)
		MutexDoc.EndReader()

		data := append([]byte(token), StringToBytes(doc)...)
		conn.Write(append([]byte{1}, data...))
	} else {
		conn.Write([]byte{0})
	}
	conn.Close()
}
