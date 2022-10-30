package ast

import (
	"fmt"

	"github.com/ReanGD/runify/server/system"
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
	dctx apd.Context
	cond apd.Condition
}

func NewAstContext(dctx apd.Context) *AstContext {
	return &AstContext{
		dctx: dctx,
		cond: apd.Condition(0),
	}
}

func NewDefaultAstContext() *AstContext {
	return NewAstContext(BaseAdpContext)
}

func toAstContext(v interface{}) (*AstContext, error) {
	if result, ok := v.(*AstContext); ok {
		return result, nil
	}

	return nil, fmt.Errorf("invalid type of parser context; expected *AstContext, got %T", v)
}

func (c *AstContext) Reset() {
	c.cond = apd.Condition(0)
}

func (c *AstContext) Error() (apd.Condition, system.Error) {
	cond := c.cond
	cond &= (c.dctx.Traps | TypeMismatch)

	if cond == 0 {
		return cond, system.Success
	}

	if cond&TypeMismatch != 0 {
		return cond, system.CalculatorTypeMismatch
	}

	if cond&apd.SystemOverflow != 0 || cond&apd.Overflow != 0 {
		return cond, system.CalculatorResultTooBig
	}

	// Result of apd.Rem or apd.QuoInteger is more than Precision
	if cond&apd.DivisionImpossible != 0 {
		return cond, system.CalculatorResultTooBig
	}

	// Result is NaN
	if cond&apd.InvalidOperation != 0 {
		return cond, system.CalculatorResultTooBig
	}

	if cond&apd.SystemUnderflow != 0 || cond&apd.Underflow != 0 {
		return cond, system.CalculatorResultTooSmall
	}

	if cond&apd.DivisionByZero != 0 || cond&apd.DivisionUndefined != 0 {
		return cond, system.CalculatorDivisionByZero
	}

	if cond&apd.Inexact != 0 || cond&apd.Subnormal != 0 || cond&apd.Rounded != 0 || apd.Clamped != 0 {
		return cond, system.CalculatorResultRounded
	}

	return cond, system.Success
}

func (c *AstContext) stringToValue(d *Value, x string) error {
	var err error
	_, c.cond, err = d.value.SetString(x)
	return err
}

func (c *AstContext) pos(d, x *Value) error {
	d.value.Set(&x.value)
	return nil
}

func (c *AstContext) neg(d, x *Value) error {
	var err error
	c.cond, err = c.dctx.Neg(&d.value, &x.value)
	return err
}

func (c *AstContext) add(d, x, y *Value) error {
	var err error
	c.cond, err = c.dctx.Add(&d.value, &x.value, &y.value)
	return err
}

func (c *AstContext) sub(d, x, y *Value) error {
	var err error
	c.cond, err = c.dctx.Sub(&d.value, &x.value, &y.value)
	return err
}

func (c *AstContext) mul(d, x, y *Value) error {
	var err error
	c.cond, err = c.dctx.Mul(&d.value, &x.value, &y.value)
	return err
}

func (c *AstContext) div(d, x, y *Value) error {
	var err error
	c.cond, err = c.dctx.Quo(&d.value, &x.value, &y.value)
	return err
}

func (c *AstContext) pow(d, x, y *Value) error {
	var err error
	c.cond, err = c.dctx.Pow(&d.value, &x.value, &y.value)
	return err
}
