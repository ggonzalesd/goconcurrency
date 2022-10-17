package main

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

var Document string

var Users map[string]string
var UserToken map[string]string
var TokenUser map[string]string
var EvSendChan map[string]chan PkgEvent
var EvRecvChan map[string]chan PkgEvent

var MutexToken MutexWR
var MutexUsers MutexWR
var MutexChann MutexWR
var MutexDoc MutexWR
