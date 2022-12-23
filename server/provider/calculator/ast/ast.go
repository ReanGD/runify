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

func (v *Value) Value() apd.Decimal {
	return v.value
}

func NewNumber(ctx, x interface{}) (*Value, error) {
	typedCtx, err, ok := toAstContext(ctx)
	if !ok {
		return nil, err
	}

	typedX, ok := x.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid type for expression new number; expected *token.Token, got %T", x)
	}

	dst := NewValue()
	_, cond, err := dst.value.SetString(string(typedX.Lit))
	typedCtx.set(cond, err)

	if err != nil {
		return nil, fmt.Errorf(`wrong number format: %s`, err)
	}

	return dst, nil
}

func unaryExprType(x *Value) TypeID {
	return x.typeID
}

func UnaryExpr(ctx, x interface{}, op byte) (*Value, error) {
	typedCtx, err, ok := toAstContext(ctx)
	if !ok {
		return nil, err
	}

	typedX, ok := x.(*Value)
	if !ok {
		return nil, fmt.Errorf(`invalid type for expression "%s(X)"; expected X is *Value, got %T`, string([]byte{op}), x)
	}

	typedX.typeID = unaryExprType(typedX)

	switch op {
	case '+':
		// pass
	case '-':
		typedCtx.set(typedCtx.dctx.Neg(&typedX.value, &typedX.value))
	default:
		return nil, fmt.Errorf(`unknown unary operator "%s"`, string([]byte{op}))
	}

	return typedX, nil
}

func binaryExprType(x, y *Value, op byte) (TypeID, error) {
	if x.typeID != y.typeID {
		return NoType, fmt.Errorf(`type mismatch for expression "X %s Y": %v != %v`, string([]byte{op}), x.typeID, y.typeID)
	}

	return x.typeID, nil
}

func BinaryExpr(ctx, x, y interface{}, op byte) (*Value, error) {
	typedCtx, err, ok := toAstContext(ctx)
	if !ok {
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

	typedX.typeID, err = binaryExprType(typedX, typedY, op)
	if err != nil {
		typedCtx.set(TypeMismatch, err)
		return nil, nil
	}

	switch op {
	case '+':
		typedCtx.set(typedCtx.dctx.Add(&typedX.value, &typedX.value, &typedY.value))
	case '-':
		typedCtx.set(typedCtx.dctx.Sub(&typedX.value, &typedX.value, &typedY.value))
	case '*':
		typedCtx.set(typedCtx.dctx.Mul(&typedX.value, &typedX.value, &typedY.value))
	case '/':
		typedCtx.set(typedCtx.dctx.Quo(&typedX.value, &typedX.value, &typedY.value))
	case '^':
		typedCtx.set(typedCtx.dctx.Pow(&typedX.value, &typedX.value, &typedY.value))
	default:
		return nil, fmt.Errorf(`unknown binary operator "%s"`, string([]byte{op}))
	}

	return typedX, nil
}
