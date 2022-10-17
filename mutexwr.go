package main

import "sync"

type MutexWR struct {
	readers     uint32
	wantWrite   sync.Mutex
	freeWriter  sync.Mutex
	readerMutex sync.Mutex
}

func (m *MutexWR) StartWrite() {
	m.wantWrite.Lock()
	m.freeWriter.Lock()
}

func (m *MutexWR) EndWrite() {
	m.freeWriter.Unlock()
	m.wantWrite.Unlock()
}

func (m *MutexWR) StartReader() {
	m.wantWrite.Lock()
	m.readerMutex.Lock()
	m.readers++
	if m.readers == 1 {
		m.freeWriter.Lock()
	}
	m.readerMutex.Unlock()
	m.wantWrite.Unlock()
}

func (m *MutexWR) EndReader() {
	m.readerMutex.Lock()
	m.readers--
	if m.readers == 0 {
		m.freeWriter.Unlock()
	}
	m.readerMutex.Unlock()
}
