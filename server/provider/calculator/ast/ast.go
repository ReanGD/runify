package ast

import (
	"fmt"

	"github.com/ReanGD/runify/server/provider/calculator/gocc/token"
	"github.com/cockroachdb/apd/v3"
)

type TypeID uint16

const (
	NoType TypeID = iota
	Currency
)

type Value struct {
	value  apd.Decimal
	typeID TypeID
}

func NewValue() *Value {
	return &Value{typeID: NoType}
}

func NewValueUnaryExprDeductionType(x *Value) *Value {
	return &Value{typeID: x.typeID}
}

func NewValueBinaryExprDeductionType(x, y *Value, op byte) (*Value, error) {
	if x.typeID != y.typeID {
		return nil, fmt.Errorf(`type mismatch for expression "X %s Y": %v != %v`, string([]byte{op}), x.typeID, y.typeID)
	}

	return &Value{typeID: x.typeID}, nil
}

func (v *Value) Value() apd.Decimal {
	return v.value
}

func NewNumber(ctx, x interface{}) (*Value, error) {
	typedCtx, err := toAstContext(ctx)
	if err != nil {
		return nil, err
	}
	typedX, ok := x.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid type for expression new number; expected *token.Token, got %T", x)
	}

	dst := NewValue()
	if err = typedCtx.stringToValue(dst, string(typedX.Lit)); err != nil {
		return nil, fmt.Errorf(`runtime error for expression new number: %s`, err)
	}

	return dst, nil
}

func UnaryExpr(ctx, x interface{}, op byte) (*Value, error) {
	typedCtx, err := toAstContext(ctx)
	if err != nil {
		return nil, err
	}
	typedX, ok := x.(*Value)
	if !ok {
		return nil, fmt.Errorf(`invalid type for expression "%s(X)"; expected X is *Value, got %T`, string([]byte{op}), x)
	}

	var fn func(d, x *Value) error
	switch op {
	case '+':
		fn = typedCtx.pos
	case '-':
		fn = typedCtx.neg
	default:
		return nil, fmt.Errorf(`unknown unary operator "%s"`, string([]byte{op}))
	}

	dst := NewValueUnaryExprDeductionType(typedX)
	if err = fn(dst, typedX); err != nil {
		return nil, err
	}

	return dst, nil
}

func BinaryExpr(ctx, x, y interface{}, op byte) (*Value, error) {
	typedCtx, err := toAstContext(ctx)
	if err != nil {
		return nil, err
	}
	typedX, ok := x.(*Value)
	if !ok {
		return nil, fmt.Errorf(`invalid type for expression "X %s Y"; expected X is *Value, got %T`, string([]byte{op}), x)
	}
	typedY, ok := y.(*Value)
	if !ok {
		return nil, fmt.Errorf(`invalid type for expression "X %s Y"; expected Y is *Value, got %T`, string([]byte{op}), y)
	}

	var fn func(d, x, y *Value) error
	switch op {
	case '+':
		fn = typedCtx.add
	case '-':
		fn = typedCtx.sub
	case '*':
		fn = typedCtx.mul
	case '/':
		fn = typedCtx.div
	case '^':
		fn = typedCtx.pow
	default:
		return nil, fmt.Errorf(`unknown binary operator "%s"`, string([]byte{op}))
	}

	dst, err := NewValueBinaryExprDeductionType(typedX, typedY, op)
	if err != nil {
		typedCtx.cond = TypeMismatch
		return nil, err
	}

	if err = fn(dst, typedX, typedY); err != nil {
		return nil, fmt.Errorf(`runtime error for expression "X %s Y": %s`, string([]byte{op}), err)
	}

	return dst, nil
}
