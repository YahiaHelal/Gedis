package cmd

type Command string
type HandlerFunc[Args any, Response any] func(args Args) Response

const (
	CommandPing Command = "PING"
)
