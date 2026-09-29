package handlers

import (
	"github.com/YahiaHelal/Gedis/app"
	"github.com/YahiaHelal/Gedis/cmd"
	"github.com/YahiaHelal/Gedis/errs"
	"github.com/YahiaHelal/Gedis/model"
	"github.com/YahiaHelal/Gedis/utils"
)

func HandleCommand(cmdArgs *model.CommandArgs) string {
	cmd := cmd.Command(cmdArgs.Cmd)
	handler, ok := app.Handlers.GetHandlerFunc(cmd)

	if ok {
		return handler(cmdArgs.Args)
	}
	return errs.NewUnknownCmdErr("Unknown command", cmdArgs.Cmd).Error()
}

func CmdPingHandler(args []string) string {
	if len(args) > 0 {
		return utils.EncodeBulkString(args[0])
	}
	return "+PONG\r\n"
}


// BUG: handle no args panic
func CmdEchoHandler(args []string) string {
	if len(args) == 0 {
		return errs.NewWrongArgsErr("wrong number of arguments for command", string(cmd.CommandEcho)).Error()
	}
	return utils.EncodeBulkString(args[0])
}
