package handlers

import (
	"github.com/YahiaHelal/Gedis/app"
	"github.com/YahiaHelal/Gedis/cmd"
	"github.com/YahiaHelal/Gedis/errs"
	"github.com/YahiaHelal/Gedis/model"
	"github.com/YahiaHelal/Gedis/serialize"
)

func HandleCommand(cmdArgs *model.CommandArgs) string {
	handler, ok := app.Handlers.GetHandlerFunc(cmdArgs)

	if ok {
		return handler(cmdArgs.Args)
	}
	return serialize.Error(errs.NewUnknownCmdErr("unknown command", cmdArgs.Cmd).Error())
}

func CmdPingHandler(args []string) string {
	if len(args) > 0 {
		return serialize.BulkString((args[0]))
	}
	return "+PONG\r\n"
}

func CmdEchoHandler(args []string) string {
	if len(args) == 0 {
		return serialize.Error(errs.NewWrongArgsErr("wrong number of arguments for command", string(cmd.CommandEcho)).Error())
	}
	return serialize.BulkString(args[0])
}

func CmdDocsHandler(args []string) string {
	return serialize.SimpleString("OK")
}
