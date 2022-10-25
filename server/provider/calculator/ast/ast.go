package ast

import (
	"fmt"

	"github.com/ReanGD/runify/server/provider/calculator/gocc/token"
	"github.com/shopspring/decimal"
)

type TypeID uint16

const (
	NoType TypeID = iota
	Currency
)

type Value struct {
	value  decimal.Decimal
	typeID TypeID
}

func NewValueFromToken(valToken interface{}) (*Value, error) {
	valTok, ok := valToken.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid basic literal type; expected *token.Token, got %T", valToken)
	}

	val, err := decimal.NewFromString(string(valTok.Lit))
	if err != nil {
		return nil, err
	}

	return &Value{value: val, typeID: NoType}, nil
}

func (v *Value) Value() decimal.Decimal {
	return v.value
}

func (v *Value) Neg() (*Value, error) {
	return &Value{value: v.value.Neg(), typeID: v.typeID}, nil
}

func (v *Value) Add(other *Value) (*Value, error) {
	if v.typeID != other.typeID {
		return nil, fmt.Errorf(`type mismatch for expression "X + Y": %v != %v`, v.typeID, other.typeID)
	}

	return &Value{value: v.value.Add(other.value), typeID: v.typeID}, nil
}

func (v *Value) Sub(other *Value) (*Value, error) {
	if v.typeID != other.typeID {
		return nil, fmt.Errorf(`type mismatch for expression "X - Y": %v != %v`, v.typeID, other.typeID)
	}

	return &Value{value: v.value.Sub(other.value), typeID: v.typeID}, nil
}

func (v *Value) Mul(other *Value) (*Value, error) {
	if v.typeID != other.typeID {
		return nil, fmt.Errorf(`type mismatch for expression "X * Y": %v != %v`, v.typeID, other.typeID)
	}

	return &Value{value: v.value.Mul(other.value), typeID: v.typeID}, nil
}

func (v *Value) Div(other *Value) (*Value, error) {
	if v.typeID != other.typeID {
		return nil, fmt.Errorf(`type mismatch for expression "X / Y": %v != %v`, v.typeID, other.typeID)
	}

	if other.value.Sign() == 0 {
		return nil, newDivisionByZero()
	}
	return &Value{value: v.value.Div(other.value), typeID: v.typeID}, nil
}

func UnaryExpr(x, operator interface{}) (*Value, error) {
	typedOp, ok := operator.(*token.Token)
	if !ok {
		return nil, fmt.Errorf(`invalid type for binary operator; expected Operator is *token.Token, got %T`, operator)
	}
	op := string(typedOp.Lit)
	typedX, ok := x.(*Value)
	if !ok {
		return nil, fmt.Errorf(`invalid type for "%s(X)"; expected X is *Value, got %T`, op, x)
	}

	switch op {
	case "-":
		return typedX.Neg()
	case "+":
		return typedX, nil
	default:
		return nil, fmt.Errorf(`unknown unary operator "%s"`, op)
	}
}

func BinaryExpr(x, y, operator interface{}) (*Value, error) {
	typedOp, ok := operator.(*token.Token)
	if !ok {
		return nil, fmt.Errorf(`invalid type for binary operator; expected Operator is *token.Token, got %T`, operator)
	}
	op := string(typedOp.Lit)
	typedX, ok := x.(*Value)
	if !ok {
		return nil, fmt.Errorf(`invalid type for expression "X %s Y"; expected X is *Value, got %T`, op, x)
	}
	typedY, ok := y.(*Value)
	if !ok {
		return nil, fmt.Errorf(`invalid type for expression "X %s Y"; expected Y is *Value, got %T`, op, y)
	}

	switch op {
	case "+":
		return typedX.Add(typedY)
	case "-":
		return typedX.Sub(typedY)
	case "*":
		return typedX.Mul(typedY)
	case "/":
		return typedX.Div(typedY)
	default:
		return nil, fmt.Errorf(`unknown binary operator "%s"`, op)
	}
}
