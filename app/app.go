package app

import (
	"sync"

	"github.com/YahiaHelal/Gedis/cmd"
)

type App struct {
	handlers map[cmd.Command]cmd.HandlerFunc[[]string, string]
	mu       *sync.RWMutex
}

var Handlers *App

func init() {
	Handlers = &App{handlers: make(map[cmd.Command]cmd.HandlerFunc[[]string, string]), mu: &sync.RWMutex{}}
}

func (a *App) RegisterCommandHandler(cmd cmd.Command, handler cmd.HandlerFunc[[]string, string]) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.handlers[cmd] = handler
}

func (a *App) GetHandlerFunc(cmd cmd.Command) (cmd.HandlerFunc[[]string, string], bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	handler, ok := a.handlers[cmd]
	return handler, ok
	// if !ok {
	// 	slog.Log(context.Background(), slog.LevelWarn, "Unknown handler for command: %s", cmd)
	// 	return nil
	// }
	// return handler
}
