package handlers

import (
	"github.com/YahiaHelal/Gedis/app"
	"github.com/YahiaHelal/Gedis/cmd"
)

func init() {
	registerAppCommands()
}

func registerAppCommands() {
	app.CmdDispatcher.Register(cmd.CommandPing, CmdPingHandler)
	app.CmdDispatcher.Register(cmd.CommandEcho, CmdEchoHandler)
	app.CmdDispatcher.Register(cmd.CommandDocs, CmdDocsHandler) // READS COMMAND only as the command instead of command odcs
}
