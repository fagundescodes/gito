package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	gitHashObject := flag.NewFlagSet("hash-object", flag.ExitOnError)
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
			if err := os.Mkdir(".gito/objects", 0o755); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			wd, err := os.Getwd()
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
			oid, err := hashObject(gitHashObject.Arg(0))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(oid)

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

func hashObject(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha1.Sum(data)
	oid := hex.EncodeToString(sum[:])
	if err := os.WriteFile(".gito/objects/"+oid, data, 0o644); err != nil {
		return "", err
	}
	return oid, nil
}
