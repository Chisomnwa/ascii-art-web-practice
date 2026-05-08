package main

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(text, banner string) (string, error) {
	// Make sure the user chooses the correct banner
	valid := banner == "standard" || banner == "thinkertoy" || banner == "shadow"
	if !valid {
		fmt.Println("wrong banner usage")
		return "", fmt.Errorf("invalid banner")
	}

	// Read the content of the banner file
	content, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("error reading file: ", err)
		return "", fmt.Errorf("Cannot read banner file")
	}

	data := strings.Split(string(content), "\n")

	// Now, work on the string so we can be able to manipulate it
	inputText := strings.ReplaceAll(text, "\\n", "\n")
	wordSlice := strings.Split(inputText, "\n")

	// Create a variable that will hold our final result
	var result strings.Builder

	// Loop through eah word, through each character of each word
	// And store their ascii character in response
	for i, word := range wordSlice {
		// To avoid unncessary leading blank line
		if word == "" {
			if i != 0 {
				result.WriteString("\n")
				continue
			}
		}

		for i := 0; i < len(word); i++ {
			for _, char := range word {
				asciiIndex := int(char - ' ') * 9 + 1 + i
				charIndex := data[asciiIndex]
				result.WriteString(charIndex)
			}
			result.WriteString("\n")
		}
		}
		return result.String(), nil
	}
	