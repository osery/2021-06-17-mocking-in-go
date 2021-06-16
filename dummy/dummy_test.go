package dummy

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/osery/2021-06-17-mocking-in-go/pkg/bar"
)

// START_TEST1 OMIT
func TestDummySuccess(t *testing.T) {
	// Given.
	d := &dummy{
		result: "a",
	}
	b := bar.NewBar(d)

	// When.
	actual, err := b.Bar("x", 3)

	// Then.
	require.Equal(t, 3, d.called)
	require.NoError(t, err)
	require.Equal(t, "aaa", actual)
}

// END_TEST1 OMIT

// START_TEST2 OMIT
func TestDummyFailure(t *testing.T) {
	// Given.
	expected := errors.New("expected")
	d := &dummy{
		err: expected,
	}
	b := bar.NewBar(d)

	// When.
	_, err := b.Bar("x", 3)

	// Then.
	require.Equal(t, 1, d.called)
	require.ErrorIs(t, err, expected)
	require.Contains(t, err.Error(), "fooing string 'x'")
	require.Contains(t, err.Error(), "step 0")
}

// END_TEST2 OMIT

// START_DUMMY OMIT
type dummy struct {
	called int
	result string
	err    error
}

func (d *dummy) Foo(s string) (string, error) {
	d.called++
	return d.result, d.err
}

// END_DUMMY OMIT
