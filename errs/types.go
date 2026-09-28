package errs

import "fmt"

type UnknownCmdError struct {
	cmd     string
	message string
}

func (e *UnknownCmdError) Error() string {
	return fmt.Sprintf("-ERR %s '%s'\r\n", e.message, e.cmd)
}
