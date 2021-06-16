package mocking

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/osery/2021-06-17-mocking-in-go/pkg/bar"
	"github.com/osery/2021-06-17-mocking-in-go/pkg/testing/generated/mockgen/mocks"
)

// START_TEST1 OMIT
func TestMockSuccess(t *testing.T) {
	// Given.
	c := gomock.NewController(t) // HL
	defer c.Finish()             // HL
	m := mocks.NewMockFoo(c)     // HL
	b := bar.NewBar(m)           // HL

	m.EXPECT(). // HL
		Foo("x"). // HL
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
	c := gomock.NewController(t)
	defer c.Finish()
	m := mocks.NewMockFoo(c)
	b := bar.NewBar(m)

	m.EXPECT(). // HL
		Foo("x"). // HL
		Return("", expected) // HL

	// When.
	_, err := b.Bar("x", 3)

	// Then.
	require.ErrorIs(t, err, expected)
	require.Contains(t, err.Error(), "fooing string 'x'")
	require.Contains(t, err.Error(), "step 0")
}

// END_TEST2 OMIT

// START_TEST3 OMIT
func TestMockCases(t *testing.T) {
	// Given.
	expected := errors.New("expected")
	c := gomock.NewController(t)
	defer c.Finish()
	m := mocks.NewMockFoo(c)
	b := bar.NewBar(m)

	gomock.InOrder( // HL
		m.EXPECT().Foo("x").Return("a", nil), // HL
		m.EXPECT().Foo("x").Return("b", nil), // HL
		m.EXPECT().Foo("x").Return("c", nil), // HL
	) // HL

	// When.
	actual, err := b.Bar("x", 3)

	// Then.
	require.NoError(t, err, expected)
	require.Equal(t, "abc", actual)
}

// END_TEST3 OMIT

// START_TEST4 OMIT
func TestMockMatching(t *testing.T) {
	// Given.
	expected := errors.New("expected")
	c := gomock.NewController(t)
	defer c.Finish()
	m := mocks.NewMockFoo(c)
	b := bar.NewBar(m)

	m.EXPECT(). // HL
		Foo(gomock.Any()). // HL
		Times(3). // HL
		DoAndReturn(func(s string) (string, error) { // HL
			return s, nil // HL
	}) // HL

	// When.
	actual, err := b.Bar("x", 3)

	// Then.
	require.NoError(t, err, expected)
	require.Equal(t, "xxx", actual)
}

// END_TEST4 OMIT
