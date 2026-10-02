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
	app.Handlers.RegisterCommandHandler(cmd.CommandEcho, CmdEchoHandler)
	app.Handlers.RegisterCommandHandler(cmd.CommandDocs, CmdDocsHandler) // READS COMMAND only as the command instead of command odcs
}
