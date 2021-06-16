package suites

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/osery/2021-06-17-mocking-in-go/pkg/bar"
	"github.com/osery/2021-06-17-mocking-in-go/pkg/testing/generated/mockery/mocks"
)

// START_TEST1 OMIT
func TestMockSuccess(t *testing.T) {
	// Given.
	m := new(mocks.Foo)           // HL
	defer m.AssertExpectations(t) // HL
	b := bar.NewBar(m)

	m.On("Foo", "x"). // HL
		Times(3). // HL
		Return("a", nil) // HL

	// When.
	actual, err := b.Bar("x", 3)

	// Then.
	require.NoError(t, err)
	require.Equal(t, "aaa", actual)
}

// END_TEST1 OMIT

// START_TEST2 OMIT
func TestMockFailure(t *testing.T) {
	// Given.
	expected := errors.New("expected")
	m := new(mocks.Foo)           // HL
	defer m.AssertExpectations(t) // HL
	b := bar.NewBar(m)

	m.On("Foo", "x"). // HL
		Return("", expected) // HL

	// When.
	_, err := b.Bar("x", 3)

	// Then.
	require.ErrorIs(t, err, expected)
	require.Contains(t, err.Error(), "fooing string 'x'")
	require.Contains(t, err.Error(), "step 0")
}

// END_TEST2 OMIT
