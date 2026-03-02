package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/TomislavRigo/frep/internal/scanner"
)

func main() {
	pattern := os.Args[1]
	rawPath := os.Args[2]

	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	path, err := filepath.Abs(filepath.Clean(rawPath))
	if err != nil {
		panic(err)
	}

	scanner.ScanAndMatch(path, pattern)
}
