package main

import (
	"fmt"
	"os"
	raosay "rao-says/bot"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: raosays <message>")
		os.Exit(1)
	}

	message := strings.Join(os.Args[1:], "")

	bot := raosay.Random()

	speechBubble := createSpeechBubble(message)

	fmt.Println(speechBubble)
	fmt.Println(bot)
}

func createSpeechBubble(message string) string {
	lines := strings.Split(message, "\n")

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	border := "+" + strings.Repeat("-", maxLen+2) + "+"

	speechBubble := border + "\n"

	for _, line := range lines {
		speechBubble += fmt.Sprintf("| %-*s |\n", maxLen, line)
	}

	speechBubble += border

	return speechBubble
}
