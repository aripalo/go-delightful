package ruler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		name     string
		char     string
		count    int
		received string
	}{
		{
			name:     "zero",
			char:     "=",
			count:    0,
			received: "",
		},
		{
			name:     "empty",
			char:     "",
			count:    3,
			received: "",
		},
		{
			name:     "basic",
			char:     "=",
			count:    3,
			received: "===",
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			received := repeat(test.char, test.count)
			assert.Equal(t, test.received, received)
		})
	}
}
