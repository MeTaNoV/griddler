package griddler

import (
	"errors"
	"fmt"
)

type ParseError struct {
	line int
	err  error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Error line %d: %s", e.line, e.err)
}

var (
	ErrInvalidGridSizeFormat = errors.New("invalid format for first line")
	ErrInvalidGridSizeValue  = errors.New("invalid value for griddler size")
	ErrMissingSemiColon      = errors.New("missing semicolon to delimit line info and values")
	ErrInvalidIntValue       = errors.New("invalid integer for value(s)")
	ErrInvalidIntLine        = errors.New("invalid integer for line info")
	ErrInvalidTokenLine      = errors.New("invalid starting token for line info")
	ErrTooManyLine           = errors.New("too many line compared to the size specified")
)

type SolveError struct {
	s   *Square
	err error
}

var (
	ErrOverridingValue = errors.New("attempt to override an existing different value")
)

func (e *SolveError) Error() string {
	return fmt.Sprintf("Error on square (%d,%d): %s", e.s.x+1, e.s.y+1, e.err)
}
