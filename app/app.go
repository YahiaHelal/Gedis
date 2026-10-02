package app

import (
	"sync"

	"github.com/YahiaHelal/Gedis/cmd"
	"github.com/YahiaHelal/Gedis/model"
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

func (a *App) GetHandlerFunc(cmdArgs *model.CommandArgs) (cmd.HandlerFunc[[]string, string], bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if handler, ok := a.handlers[cmd.Command(cmdArgs.Cmd)]; ok {
		return handler, ok
	}
	return nil, false
}

func (a *App) HasHandler(cmd cmd.Command) bool {
	_, ok := a.handlers[cmd]
	return ok
}
