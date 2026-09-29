package handlers

import (
	"github.com/YahiaHelal/Gedis/app"
	"github.com/YahiaHelal/Gedis/cmd"
)

func init() {
	registerAppCommands()
}

func registerAppCommands() {
	app.Handlers.RegisterCommandHandler(cmd.CommandPing, CmdPingHandler)
}
