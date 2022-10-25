package ast

import (
	"github.com/ReanGD/runify/server/system"
)

type Error struct {
	s       string
	errCode system.Error
}

func (e *Error) Error() string {
	return e.s
}

func (e *Error) Code() system.Error {
	return e.errCode
}

func newDivisionByZero() error {
	return &Error{"division by zero", system.CalculatorDivisionByZero}
}
