package cmd

import (
	"fmt"
)

// help command. show help message + list of available commands
func runDefault() int {
	fmt.Println("sage")
	fmt.Println()
	fmt.Println("AI-powered Git intelligence assistant.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  sage <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  commit:   auto-generate commit message from staged changes")
	fmt.Println("  exec:     generate and execute git/gh commands from natural language")
	fmt.Println("  config:   show and set configuration values")
	fmt.Println("  version:  show version information")
	return 0
}
