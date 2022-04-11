package ruler

import (
	"golang.org/x/term"
)

// getWidth returns the terminal column count or a fallback value
func getWidth(fallback int) int {
	if term.IsTerminal(0) {
		width, _, err := term.GetSize(0)
		if err == nil {
			return width
		}
	}
	return fallback
}

// repeat given character times count
func repeat(char string, count int) string {
	ruler := ""
	for i := 0; i < count; i++ {
		ruler += char
	}
	return ruler
}

// Generate a horizontal ruler.
func Generate(char string) string {
	width := getWidth(16)
	return repeat("=", width)
}
