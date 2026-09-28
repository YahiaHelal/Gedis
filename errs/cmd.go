package errs

func NewUnknownCmdErr(message string, cmd string) *UnknownCmdError {
	return &UnknownCmdError{message: message, cmd: cmd}
}
