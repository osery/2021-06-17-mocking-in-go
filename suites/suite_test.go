package suites

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/osery/2021-06-17-mocking-in-go/pkg/bar"
	"github.com/osery/2021-06-17-mocking-in-go/pkg/testing/generated/mockgen/mocks"
)

// START_SUITE OMIT
func TestSuite(t *testing.T) {
	suite.Run(t, new(mockSuite))
}

type mockSuite struct {
	suite.Suite

	c *gomock.Controller // HL
	m *mocks.MockFoo     // HL
}

func (s *mockSuite) SetupTest() {
	s.c = gomock.NewController(s.T()) // HL
	s.m = mocks.NewMockFoo(s.c)       // HL
}

func (s *mockSuite) TearDownTest() {
	s.c.Finish() // HL
}

// END_SUITE OMIT

// START_TEST1 OMIT
func (s *mockSuite) TestMockSuccess() {
	// Given.
	b := bar.NewBar(s.m)

	s.m.EXPECT(). // HL
		Foo("x"). // HL
		Times(3). // HL
		Return("a", nil) // HL

	// When.
	actual, err := b.Bar("x", 3)

	// Then.
	s.Require().NoError(err)
	s.Require().Equal("aaa", actual)
}

// END_TEST1 OMIT

// START_TEST2 OMIT
func (s *mockSuite) TestMockFailure() {
	// Given.
	expected := errors.New("expected")
	b := bar.NewBar(s.m)

	s.m.EXPECT(). // HL
		Foo("x"). // HL
		Return("", expected) // HL

	// When.
	_, err := b.Bar("x", 3)

	// Then.
	s.Require().ErrorIs(err, expected)
	s.Require().Contains(err.Error(), "fooing string 'x'")
	s.Require().Contains(err.Error(), "step 0")
}

// END_TEST2 OMIT

// START_TEST3 OMIT
func (s *mockSuite) TestMockCases() {
	// Given.
	expected := errors.New("expected")
	b := bar.NewBar(s.m)

	gomock.InOrder( // HL
		s.m.EXPECT().Foo("x").Return("a", nil), // HL
		s.m.EXPECT().Foo("x").Return("b", nil), // HL
		s.m.EXPECT().Foo("x").Return("c", nil), // HL
	) // HL

	// When.
	actual, err := b.Bar("x", 3)

	// Then.
	s.Require().NoError(err, expected)
	s.Require().Equal("abc", actual)
}

// END_TEST3 OMIT

// START_TEST4 OMIT
func (s *mockSuite) TestMockMatching() {
	// Given.
	expected := errors.New("expected")
	b := bar.NewBar(s.m)

	s.m.EXPECT(). // HL
		Foo(gomock.Any()). // HL
		Times(3). // HL
		DoAndReturn(func(s string) (string, error) { // HL
			return s, nil // HL
		}) // HL

	// When.
	actual, err := b.Bar("x", 3)

	// Then.
	s.Require().NoError(err, expected)
	s.Require().Equal("xxx", actual)
}

// END_TEST4 OMIT
