package serialize

import "fmt"

func SimpleString(s string) string {
	return fmt.Sprintf("+%s\r\n", s)
}

func Error(msg string) string {
	return fmt.Sprintf("-ERR %s\r\n", msg)
}

func Integer(num string) string {
	return fmt.Sprintf(":%s\r\n", num)
}

func BulkString(s string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
}

func Null() string {
	return "-1\r\n"
}
