package cmd

type Command string
type HandlerFunc[Args any, Response any] func(args Args) Response


// TODO: docs for each command to run along with COMMAND DOCS ECHO/PING/...
const (
	CommandPing Command = "PING"
	CommandEcho Command = "ECHO"
	CommandDocs Command = "COMMAND DOCS"
)
