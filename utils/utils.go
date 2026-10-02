package utils

func RemoveAtIndex(idx int, args []string) []string {
	args = append(args[:idx], args[idx+1:]...)
	return args
}
