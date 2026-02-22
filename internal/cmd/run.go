package cmd

import (
	"fmt"
)

// help command. show help message + list of available commands
func runDefault() int {
	fmt.Println("git-sage")
	fmt.Println()
	fmt.Println("Git intelligence CLI tool.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  git-sage <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  commit:   auto-generate commit message from staged changes")
	fmt.Println("  config:   show and set configuration values")
	fmt.Println("  version:  show version information")
	return 0
}
