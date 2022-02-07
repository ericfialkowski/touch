package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <one or more files to create or update times>\n", os.Args[0])
		os.Exit(1)
	}
	for _, arg := range os.Args[1:] {
		_, err := os.Open(arg)
		if errors.Is(err, os.ErrNotExist) {
			// create empty file
			_, err := os.Create(arg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			currentTime := time.Now().Local()
			err := os.Chtimes(arg, currentTime, currentTime)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
