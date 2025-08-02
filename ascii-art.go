package main

import (
	"fmt"
	"os"
	"strings"
)

func run(input, banner string) (string, error) {
	bannerFile := "banners/standard.txt"
	switch banner {
	case "standard", "thinkertoy", "shadow":
		bannerFile = "banners/" + banner + ".txt"
	default:
		return "", fmt.Errorf("invalid banner specified: %s", banner)
	}

	data, err := os.ReadFile(bannerFile)
	if err != nil {
		return "", fmt.Errorf("failed to read banner file: %v", err)
	}

	charMap, err := buildCharMap(data)
	if err != nil {
		return "", err
	}

	return printAsciiArt(input, charMap), nil
}

func buildCharMap(data []byte) (map[rune][]string, error) {
	lines := strings.Split(string(data), "\n")
	chars := make(map[rune][]string)

	i := 0
	r := rune(32)
	for r < 127 {
		chars[r] = lines[i : i+9]
		i += 9
		r++
	}

	if r != 127 {
		return nil, fmt.Errorf("incomplete banner file, missing characters")
	}

	return chars, nil
}

func printAsciiArt(input string, chars map[rune][]string) string {
	lines := strings.Split(input, "\r\n")

	art := ""
	for _, lineStr := range lines {

		if lineStr == "" {
			art += "\n"
			continue
		}

		for line := 1; line <= 8; line++ {

			for _, ch := range lineStr {
				if (ch < 32 || ch > 126) && (ch != 13 && ch != 10) {
					return "unsupported character"
				}
				artLines := chars[ch]
				art += (artLines[line])
			}
			art += "\n"
		}
	}
	return art
}
