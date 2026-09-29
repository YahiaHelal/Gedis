package parser

import (
	"strings"

	"github.com/YahiaHelal/Gedis/model"
)

// TODO: use stack to evaluate other than quotes arguments, like () [] {}
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
	return &model.CommandArgs{Cmd: args[0], Args: args[1:]}
}
