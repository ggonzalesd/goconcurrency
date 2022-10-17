package main

import (
	"encoding/binary"
)

func StringToBytes(s string) []byte {
	buff := make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, uint32(len(s)))
	return append(buff, []byte(s)...)
}

func BytesToString(bs []byte) (s string, l uint32) {
	ss := binary.LittleEndian.Uint32(bs[:4])
	l = 4 + ss
	s = string(bs[4:l])
	return
}

func PkgEventToBytes(pkge PkgEvent) []byte {
	buff := make([]byte, 8)
	binary.LittleEndian.PutUint32(buff[:4], uint32(pkge.start))
	binary.LittleEndian.PutUint32(buff[4:], uint32(pkge.end))
	s := StringToBytes(pkge.value)
	return append(buff, s...)
}

func BytesToPkgEvent(bs []byte) (pkge PkgEvent, l uint32) {
	pkge.start = binary.LittleEndian.Uint32(bs[:4])
	pkge.end = binary.LittleEndian.Uint32(bs[4:8])
	pkge.value, l = BytesToString(bs[8:])
	l += 8
	return
}

func UserToBytes(u User) []byte {
	return append(StringToBytes(u.name), StringToBytes(u.pass)...)
}

func BytesToUser(bs []byte) (u User, l uint32) {
	s1, l1 := BytesToString(bs)
	s2, l2 := BytesToString(bs[l1:])
	l = l1 + l2
	u.name = s1
	u.pass = s2
	return
}
