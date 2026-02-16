package base

import (
	"fmt"
	"os"
)

func WriteTree(directory string) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		full := directory + "/" + entry.Name()

		if entry.Type().IsRegular() {
			fmt.Println(full)
			continue
		}

		if entry.IsDir() {
			if err := WriteTree(full); err != nil {
				return err
			}
		}
	}

	return nil
}
