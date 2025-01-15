package strings

import (
	"fmt"
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - indexes: expected (%d) <> found (%d)"

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Coder is great", "Coder", 3}, // change  second atribute for "e" and last atribute for 11 to see error message
		{"", "", 0},
		{"Hey", "hey", -1},
		{"yuri", "y", 0},
	}
	fmt.Println(tests)

	for _, test := range tests {
		t.Logf("Excecution: %v", test)
		actual := strings.Index(test.text, test.part)

		if actual != test.expected {
			t.Errorf(msgIndex,
				test.text,
				test.part,
				test.expected,
				actual)
		}
	}
}
