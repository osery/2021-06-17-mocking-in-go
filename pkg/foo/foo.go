package foo

//go:generate go run github.com/golang/mock/mockgen -destination ../testing/generated/mockgen/mocks/foo.go -package mocks . Foo
//go:generate go run github.com/vektra/mockery/v2 --name Foo --output=../testing/generated/mockery/mocks

// START_FOO OMIT
type Foo interface {
	Foo(s string) (string, error) // HL
}

// END_FOO OMIT
