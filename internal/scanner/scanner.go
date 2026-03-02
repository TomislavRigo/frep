package scanner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	matcher "github.com/TomislavRigo/frep/internal/search"
)

func scanAndMathc(path, pattern string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	directories := []string{}
	for _, entry := range entries {
		filePath := filepath.Join(path, entry.Name())
		if entry.IsDir() && entry.Name()[0] != '.' {
			directories = append(directories, filePath)
		}

		if entry.Name()[0] == '.' {
			continue
		}

		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)
		headerPrinted := false
		lineNo := 1
		for scanner.Scan() {
			line := scanner.Text()
			for _, word := range strings.Split(line, " ") {
				if strings.TrimSpace(word) == "" {
					continue
				}
				distance := matcher.CalculateDistance(strings.TrimSpace(word), pattern)
				if distance >= 0.7 {
					if !headerPrinted {
						fmt.Println()
						fmt.Println(filePath)
						headerPrinted = true
					}

					bound := min(len(line), 120)
					fmt.Printf("ln:%d: %s\n", lineNo, line[:bound])
					break
				}
			}
			lineNo += 1
		}
	}

	for _, dir := range directories {
		scanAndMathc(dir, pattern)
	}
}

func ScanAndMatch(path, pattern string) {
	scanAndMathc(path, pattern)
}
