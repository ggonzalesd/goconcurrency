package main

import (
	"fmt"
	"strings"
)

func editing_pkge(user string, pkge PkgEvent) {
	MutexDoc.StartReader()
	l := uint32(len(Document))
	if pkge.start > l {
		pkge.start = l
	}
	if pkge.end > l {
		pkge.end = l
	}
	MutexDoc.EndReader()

	MutexDoc.StartWrite()
	Document = Document[:pkge.start] + pkge.value + Document[pkge.end:]
	fmt.Println()
	fmt.Println(Document)
	MutexDoc.EndWrite()

	for userRcv, cRcv := range EvSendChan {
		if strings.Compare(userRcv, user) != 0 {
			cRcv <- pkge
		}
	}

}

func editing() {
	for {
		MutexChann.StartReader()
		for userRcv, cRcv := range EvRecvChan {
			select {
			case keyMsg, ok := <-cRcv:
				if ok {
					editing_pkge(userRcv, keyMsg)
				}
			default:
				// Nothing
			}
		}
		MutexChann.EndReader()
	}
}
