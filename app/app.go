package app

import (
	"github.com/YahiaHelal/Gedis/cmd"
	"github.com/YahiaHelal/Gedis/dispatcher"
)

var CmdDispatcher *dispatcher.Dispatcher[cmd.Command, cmd.HandlerFunc[[]string, string]]

func init() {
	CmdDispatcher = dispatcher.NewDispatcher[cmd.Command, cmd.HandlerFunc[[]string, string]]()
}


