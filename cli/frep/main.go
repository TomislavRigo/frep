package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/TomislavRigo/frep/internal/scanner"
)

func main() {
	argLen := len(os.Args)
	if argLen < 2 {
		fmt.Printf("Missing pattern")
		os.Exit(1)
	}
	pattern := os.Args[1]
	rawPath := "."
	if argLen == 3 {
		rawPath = os.Args[2]
	}

	path, err := filepath.Abs(filepath.Clean(rawPath))
	if err != nil {
		fmt.Println("Path invalid")
		os.Exit(1)
	}

	scanner.ScanAndMatch(path, pattern)
	os.Exit(0)
}
