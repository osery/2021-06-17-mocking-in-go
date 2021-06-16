package parameterized

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// START OMIT
func TestParameterized(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{"empty", "", []string{""}},
		{"two", "x,y", []string{"x", "y"}},
		{"middle empty", "x,,y", []string{"x", "", "y"}},
		{"trailing empty", "x,y,", []string{"x", "y", ""}},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := strings.Split(tt.input, ",")
			require.Equal(t, tt.expected, actual) // HL
		})
	}
}

// END OMIT
