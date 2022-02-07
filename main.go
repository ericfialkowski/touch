package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	if runtime.GOOS == "windows" {
		if len(os.Args) == 1 {
			fmt.Printf("Usage: %s <one or more files to create or update times>\n", os.Args[0])
			os.Exit(1)
		}
		for _, arg := range os.Args[1:] {
			_, err := os.Open(arg)
			if os.IsNotExist(err) {
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
	} else {
		fmt.Println("You should already have a native touch command. Use it.")
		os.Exit(2)
	}
}
