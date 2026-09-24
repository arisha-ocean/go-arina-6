package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert detects the input format and converts text to Morse code or Morse code to text.
func Convert(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("input is empty")
	}

	cleaned := strings.ReplaceAll(input, ".", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	cleaned = strings.ReplaceAll(cleaned, " ", "")

	if cleaned == "" {
		codes := strings.Fields(input)

		for _, code := range codes {
			if morse.MorseToRune(code) == 0 {
				return "", fmt.Errorf("invalid morse code: %s", code)
			}
		}

		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}
