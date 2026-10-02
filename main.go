package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/YahiaHelal/Gedis/handlers"
	"github.com/YahiaHelal/Gedis/parser"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		args := parser.ParseArgs(line)
		response := handlers.HandleCommand(args)
		fmt.Print(response)
	}
}
