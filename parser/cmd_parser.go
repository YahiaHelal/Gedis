package parser

import (
	"strings"

	"github.com/YahiaHelal/Gedis/app"
	"github.com/YahiaHelal/Gedis/cmd"
	"github.com/YahiaHelal/Gedis/model"
	"github.com/YahiaHelal/Gedis/utils"
)

// TODO: use stack to evaluate other than quotes arguments, like () [] {}
// TODO: each cmd should have it's own way of parsing arguments, multiple words commands for example
func ParseArgs(line string) *model.CommandArgs {
	var args []string
	var current strings.Builder // single buffer
	inQuotes := false
	for _, ch := range line {
		switch {
		case ch == '"' && !inQuotes:
			inQuotes = true
		case ch == '"' && inQuotes:
			inQuotes = false
		case ch == ' ' && !inQuotes:
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(ch)
		}
	}
	if current.Len() > 0 {
		args = append(args, current.String())
	}
	cmdArgs := model.CommandArgs{Cmd: args[0], Args: args[1:]}
	tryConstruct(&cmdArgs)
	return &cmdArgs
}

// shortest to match
func tryConstruct(cmdArgs *model.CommandArgs) {
	var command strings.Builder
	var removedIdx []int
	command.WriteString(strings.ToUpper(cmdArgs.Cmd))

	for idx, arg := range cmdArgs.Args {
		trialCmd := cmd.Command((command.String()))
		if app.CmdDispatcher.HasHandler(trialCmd) {
			break
		}

		command.WriteString(" ")
		command.WriteString(strings.ToUpper(arg))
		removedIdx = append(removedIdx, idx)
	}

	for _, i := range removedIdx {
		cmdArgs.Args = utils.RemoveAtIndex(i, cmdArgs.Args)
	}
	cmdArgs.Cmd = command.String()
}
