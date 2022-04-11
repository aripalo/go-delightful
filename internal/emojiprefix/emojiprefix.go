package emojiprefix

import (
	"fmt"
	"unicode/utf8"

	"github.com/enescakir/emoji"
)

// defaultPadding determines the amount of pad characters
const defaultPadding int = 3

// determinePad resolves the correct amount of characters needed for
// padding to achieve (mostly) vertically matching "columns" of emoji prefixes.
func determinePad(prefix emoji.Emoji) int {
	runeCount := utf8.RuneCountInString(string(prefix))
	padding := defaultPadding - (runeCount * 2)
	if padding < 0 {
		return 0
	}
	return padding
}

// padRight pads string with spaces.
func padRight(prefix emoji.Emoji) string {
	return fmt.Sprintf("%-*s", determinePad(prefix), string(prefix))
}

// Format is responsible for prefixing a string with emoji.
// Empty string for emoji prefix means the emoji prefix is ignored.
// Boolean "enable" controls if the emoji prefix is set or not.
func Format(enable bool, prefix emoji.Emoji, rendered string) string {
	// empty emoji means we'll skip it
	if string(prefix) == "" {
		return rendered
	}

	// format the string with emoji prefix if emoji mode is on
	if enable {
		return fmt.Sprintf("%s %s", padRight(prefix), rendered)
	}

	// otherwise just return the string
	return rendered
}
