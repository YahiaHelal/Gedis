package handlers

import (
	"fmt"

	"github.com/YahiaHelal/Gedis/errs"
	"github.com/YahiaHelal/Gedis/model"
)

func HandleCommand(cmdArgs *model.CommandArgs) string {
	cmd := Command(cmdArgs.Cmd)
	if handler, ok := appCommands.handlers[cmd]; ok {
		return handler(cmdArgs.Args)
	}
	return errs.NewUnknownCmdErr("Unknown command", cmdArgs.Cmd).Error()

	// switch cmd {
	// case "PING":
	// 	// TODO: Return "+PONG\r\n" for no args
	// 	// TODO: Return bulk string for PING <message>
	// }
}

func CmdPingHandler(args []string) string {
	if len(args) > 0 {
		return fmt.Sprintf("$%d\r\n%s\r\n", len(CommandPing), args[0])
	}
	return "+PONG\r\n"
}
