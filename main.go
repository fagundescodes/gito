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
			if err := os.Mkdir(".gito", 0o755); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Initialized empty Gito repository in %s/.gito\n", wd)
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
