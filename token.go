package main

import "math/rand"

func generateToken() string {
	var token []rune
	for i := 0; i < 32; i++ {
		token = append(token, rune(uint32('a')+rand.Uint32()%(1+uint32('z')-uint32('a'))))
	}
	return string(token)
}
