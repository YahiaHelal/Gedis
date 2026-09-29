package errs

import "fmt"

type UnknownCmdErr struct {
	cmd     string
	message string
}

func (e *UnknownCmdErr) Error() string {
	return fmt.Sprintf("-ERR %s '%s'\r\n", e.message, e.cmd)
}

type WrongArgsErr struct {
	cmd     string
	message string
}

func (e *WrongArgsErr) Error() string {
	return fmt.Sprintf("-ERR %s '%s'\r\n", e.message, e.cmd)
}
