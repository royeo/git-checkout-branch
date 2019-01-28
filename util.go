package main

import "strings"

func isFlag(args string) bool {
	return strings.HasPrefix(args, "-")
}
