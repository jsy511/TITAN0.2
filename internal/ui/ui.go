package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Red       = "\033[31m"
	BrightRed = "\033[91m"
	Dim       = "\033[2m"
	Reset     = "\033[0m"
)

func Clear() {
	fmt.Print("\033[2J\033[H")
}

func Logo() {
	fmt.Println(BrightRed + `
  010101010101010101010101010101010101
  101010101010101010101010101010101010
  010111111101011111110101111111010101
  101100000101010000010101000001010101
  010111111101011111110101111111010101
  101100000101010000010101000001010101
  010111111101011111110101111111010101
  101010101010101010101010101010101010
  010101010101010101010101010101010101
` + Reset)

	fmt.Println(BrightRed + "                    TITAN 0.2" + Reset)
	fmt.Println(Dim + "             SECURITY RESEARCH PLATFORM" + Reset)
}

func Header(title string) {
	fmt.Println()
	fmt.Printf("%s[ %s ]%s\n", BrightRed, strings.ToUpper(title), Reset)
	fmt.Println(Red + "----------------------------------------" + Reset)
}

func Prompt() string {
	fmt.Print(BrightRed + "\nTITAN > " + Reset)

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(input)
}

func Pause() {
	fmt.Print(Dim + "\nPress ENTER to continue..." + Reset)
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func Error(message string) {
	fmt.Printf("%s[ERROR] %s%s\n", Red, message, Reset)
}

func Success(message string) {
	fmt.Printf("%s[OK] %s%s\n", BrightRed, message, Reset)
}