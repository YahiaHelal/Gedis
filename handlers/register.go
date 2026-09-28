package handlers

import (
	"sync"
)

var appCommands *app

func init() {
	appCommands = &app{handlers: make(map[Command]HandlerFunc[[]string, string]) ,mu: &sync.RWMutex{}}
	registerAppCommands()
}

func registerAppCommands() {
	registerCommandHandler(CommandPing, CmdPingHandler)
}

func registerCommandHandler(cmd Command, handler HandlerFunc[[]string, string]) {
	appCommands.mu.Lock()
	defer appCommands.mu.Unlock()
	appCommands.handlers[cmd] = handler
}

// func GetHandlerFunc(cmd Command) HandlerFunc[[]string, string] {
// 	appCommands.mu.RLock()
// 	defer appCommands.mu.RUnlock()
// 	handler, ok := appCommands.handlers[cmd]
// 	if !ok {
// 		slog.Log(context.Background(), slog.LevelWarn, "Unknown handler for command: %s", cmd)
// 		return nil
// 	}
// 	return handler
// }
