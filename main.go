package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fagundescodes/gito/internal/base"
	"github.com/fagundescodes/gito/internal/data"
)

func main() {
	gitHashObject := flag.NewFlagSet("hash-object", flag.ExitOnError)
	gitCatFile := flag.NewFlagSet("cat-file", flag.ExitOnError)
	gitWriteTree := flag.NewFlagSet("write-tree", flag.ExitOnError)

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
			wd, err := data.Init()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Initialized empty Gito repository in %s/.gito\n", wd)
		case "hash-object":
			gitHashObject.Parse(os.Args[3:])
			if gitHashObject.NArg() < 1 {
				fmt.Println("expected a file")
				os.Exit(1)
			}

			content, err := os.ReadFile(gitHashObject.Arg(0))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			oid, err := data.HashObject(content, "blob")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(oid)
		case "cat-file":
			gitCatFile.Parse(os.Args[3:])
			if gitCatFile.NArg() < 1 {
				fmt.Println("Expected an object")
				os.Exit(1)
			}
			content, err := data.GetObject(gitCatFile.Arg(0))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Print(string(content))
		case "write-tree":
			gitWriteTree.Parse(os.Args[3:])
			if err := base.WriteTree("."); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		default:
			fmt.Printf("Unknown subcommand %s\n", os.Args[2])
			os.Exit(1)
		}

		gitCmd.Parse(os.Args[2:])
	default:
		fmt.Printf("Unknown command %s\n", os.Args[1])
		os.Exit(1)
	}
}
