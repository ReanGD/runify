package ast

import (
	"fmt"

	"github.com/ReanGD/runify/server/global"
	"github.com/cockroachdb/apd/v3"
)

const (
	DefaultAstTraps = apd.SystemOverflow |
		apd.SystemUnderflow |
		apd.Overflow |
		apd.Underflow |
		apd.DivisionUndefined |
		apd.DivisionByZero |
		apd.DivisionImpossible |
		apd.InvalidOperation
	// apd.Inexact |
	// apd.Subnormal |
	// apd.Rounded |
	// apd.Clamped
)

const (
	TypeMismatch = apd.Clamped << 1
)

var BaseAdpContext = apd.Context{
	Precision:   16,
	MaxExponent: apd.MaxExponent,
	MinExponent: apd.MinExponent,
	Traps:       DefaultAstTraps,
}

type AstContext struct {
	dctx      apd.Context
	cond      apd.Condition
	fullTraps apd.Condition
}

func NewAstContext(dctx apd.Context) *AstContext {
	return &AstContext{
		dctx:      dctx,
		cond:      apd.Condition(0),
		fullTraps: dctx.Traps | TypeMismatch,
	}
}

func NewDefaultAstContext() *AstContext {
	return NewAstContext(BaseAdpContext)
}

func toAstContext(v interface{}) (*AstContext, error, bool) {
	if result, ok := v.(*AstContext); ok {
		return result, nil, result.isValid()
	}

	return nil, fmt.Errorf("invalid type of parser context; expected *AstContext, got %T", v), false
}

func (c *AstContext) GetApdContext() apd.Context {
	return c.dctx
}

func (c *AstContext) set(cond apd.Condition, err error) {
	c.cond |= cond
}

func (c *AstContext) isValid() bool {
	return (c.cond & c.fullTraps) == 0
}

func (c *AstContext) Reset() {
	c.cond = apd.Condition(0)
}

func (c *AstContext) Error() (apd.Condition, global.Error) {
	cond := c.cond & c.fullTraps

	if cond == 0 {
		return cond, global.Success
	}

	if cond&TypeMismatch != 0 {
		return cond, global.CalculatorTypeMismatch
	}

	if cond&apd.SystemOverflow != 0 || cond&apd.Overflow != 0 {
		return cond, global.CalculatorResultTooBig
	}

	// Result of apd.Rem or apd.QuoInteger is more than Precision
	if cond&apd.DivisionImpossible != 0 {
		return cond, global.CalculatorResultTooBig
	}

	// Result is NaN
	if cond&apd.InvalidOperation != 0 {
		return cond, global.CalculatorResultTooBig
	}

	if cond&apd.SystemUnderflow != 0 || cond&apd.Underflow != 0 {
		return cond, global.CalculatorResultTooSmall
	}

	if cond&apd.DivisionByZero != 0 || cond&apd.DivisionUndefined != 0 {
		return cond, global.CalculatorDivisionByZero
	}

	if cond&apd.Inexact != 0 || cond&apd.Subnormal != 0 || cond&apd.Rounded != 0 || apd.Clamped != 0 {
		return cond, global.CalculatorResultRounded
	}

	return cond, global.Success
}
