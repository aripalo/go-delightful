package emojiprefix

import (
	"fmt"
	"testing"

	"github.com/enescakir/emoji"
	"github.com/stretchr/testify/assert"
)

func TestDeterminePad(t *testing.T) {
	tests := []struct {
		name     string
		input    emoji.Emoji
		expected int
	}{
		{
			name:     "double rune: checkbox with check",
			input:    emoji.CheckBoxWithCheck,
			expected: 3,
		},
		{
			name:     "double rune: gear",
			input:    emoji.Gear,
			expected: 3,
		},
		{
			name:     "single rune: pile of poop",
			input:    emoji.PileOfPoo,
			expected: 1,
		},
		{
			name:     "single rune: airplane departure",
			input:    emoji.AirplaneDeparture,
			expected: 1,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := determinePad(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}
