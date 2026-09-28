package handlers

import "sync"

type Command string
type HandlerFunc[Args any, Response any] func(args Args) Response

type app struct {
	handlers map[Command]HandlerFunc[[]string, string]
	mu       *sync.RWMutex
}

const (
	CommandPing Command = "PING"
)
