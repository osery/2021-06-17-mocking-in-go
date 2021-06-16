package bar

import (
	"fmt"
	"strings"

	"github.com/osery/2021-06-17-mocking-in-go/pkg/foo"
)

// START_PUBLIC OMIT
type Bar interface {
	// Bar returns a concatenation of the string s fooed n times.
	Bar(s string, n int) (string, error) // HL
}

func NewBar(foo foo.Foo) Bar {
	return &bar{
		foo: foo,
	}
}

// END_PUBLIC OMIT

// START_IMPL OMIT
type bar struct {
	foo foo.Foo
}

func (b *bar) Bar(s string, n int) (string, error) {
	var o strings.Builder
	for i := 0; i < n; i++ {
		f, err := b.foo.Foo(s)
		if err != nil {
			return "", fmt.Errorf("fooing string '%s', step %d: %w", s, i, err)
		}
		o.WriteString(f)
	}
	return o.String(), nil
}

// END_IMPL OMIT
