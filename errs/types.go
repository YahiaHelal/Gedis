package errs

import "fmt"

type UnknownCmdErr struct {
	cmd     string
	message string
}

type WrongArgsErr struct {
	cmd     string
	message string
}

func (e *UnknownCmdErr) Error() string {
	return fmt.Sprintf("%s '%s'", e.message, e.cmd)
}

func (e *WrongArgsErr) Error() string {
	return fmt.Sprintf("%s '%s'", e.message, e.cmd)
}
