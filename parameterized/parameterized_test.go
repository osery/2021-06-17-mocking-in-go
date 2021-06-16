package parameterized

import (
	"reflect"
	"strings"
	"testing"
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
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Actual %v not equal to expected %v.", actual, tt.expected)
			}
		})
	}
}

// END OMIT
