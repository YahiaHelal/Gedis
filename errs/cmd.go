package errs

func NewUnknownCmdErr(message string, cmd string) *UnknownCmdErr {
	return &UnknownCmdErr{message: message, cmd: cmd}
}

func NewWrongArgsErr(message string, cmd string) *WrongArgsErr {
	return &WrongArgsErr{message: message, cmd: cmd}
}
