package emojiprefix

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/mattn/go-runewidth"
)

// defaultPadding determines the amount of pad characters
const maxPrefixRuneLength int = 3

// determinePad resolves the correct amount of characters needed for
// padding to achieve (mostly) vertically matching "columns" of emoji prefixes.
func determinePad(prefix emoji.Emoji) int {
	width := runewidth.StringWidth(prefix.String())
	padSize := maxPrefixRuneLength - width
	return padSize
}

// padRight pads string with spaces.
func padRight(prefix emoji.Emoji) string {
	padSize := determinePad(prefix)
	result := prefix.String()

	for i := 0; i < padSize; i++ {
		result = result + string(rune(' '))
	}
	return result
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
