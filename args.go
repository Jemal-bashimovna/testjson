package main

import (
	"fmt"
	"os"
)

func args() {
	if len(os.Args) < 2 {
		fmt.Println("provide arguments")
	}

	for i, arg := range os.Args {
		fmt.Printf("arguments %d => %s\n", i, arg)
	}
}
