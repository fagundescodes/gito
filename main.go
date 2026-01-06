package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	gitCmd := flag.NewFlagSet("gito", flag.ExitOnError)
	gitInit := flag.NewFlagSet("init", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Expected at last one command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "gito":
		if len(os.Args) < 3 {
			fmt.Println("Expected a command")
			os.Exit(1)
		}

		switch os.Args[2] {
		case "init":
			gitInit.Parse(os.Args[3:])
			fmt.Println("Initialized empty Gito repository")
		default:
			fmt.Printf("Unknown subcommand %s\n", os.Args[2])
			os.Exit(1)
		}

		gitCmd.Parse(os.Args[2:])
	default:
		fmt.Printf("Unknown command %s\n", os.Args[2])
		os.Exit(1)
	}
}
