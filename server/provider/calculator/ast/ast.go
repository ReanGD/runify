package ast

import (
	"fmt"
	"strconv"

	"github.com/ReanGD/runify/server/provider/calculator/gocc/token"
)

type TypeID uint16

const (
	NoType TypeID = iota
	Currency
)

type Value struct {
	value  int64
	typeID TypeID
}

func NewValueFromToken(valToken interface{}) (*Value, error) {
	valTok, ok := valToken.(*token.Token)
	if !ok {
		return nil, fmt.Errorf("invalid basic literal type; expected *token.Token, got %T", valToken)
	}

	val, err := strconv.ParseInt(string(valTok.Lit), 10, 64)
	if err != nil {
		return nil, err
	}

	return &Value{value: val, typeID: NoType}, nil
}

func CastValue(val interface{}, op string) (*Value, error) {
	if typedVal, ok := val.(*Value); ok {
		return typedVal, nil
	}

	return nil, fmt.Errorf(`invalid type for "%s"; expected *Value, got %T`, op, val)
}

func (v *Value) Value() int64 {
	return v.value
}

func (v *Value) Neg() (*Value, error) {
	return &Value{value: -v.value, typeID: v.typeID}, nil
}

func (v *Value) Add(other *Value) (*Value, error) {
	if v.typeID != other.typeID {
		return nil, fmt.Errorf(`type mismatch for "+": %v != %v`, v.typeID, other.typeID)
	}

	return &Value{value: v.value + other.value, typeID: v.typeID}, nil
}

func (v *Value) Sub(other *Value) (*Value, error) {
	if v.typeID != other.typeID {
		return nil, fmt.Errorf(`type mismatch for "-": %v != %v`, v.typeID, other.typeID)
	}

	return &Value{value: v.value - other.value, typeID: v.typeID}, nil
}

func (v *Value) Mul(other *Value) (*Value, error) {
	if v.typeID != other.typeID {
		return nil, fmt.Errorf(`type mismatch for "*": %v != %v`, v.typeID, other.typeID)
	}

	return &Value{value: v.value * other.value, typeID: v.typeID}, nil
}

func (v *Value) Div(other *Value) (*Value, error) {
	if v.typeID != other.typeID {
		return nil, fmt.Errorf(`type mismatch for "/": %v != %v`, v.typeID, other.typeID)
	}

	return &Value{value: v.value / other.value, typeID: v.typeID}, nil
}

func Neg(a interface{}) (*Value, error) {
	aVal, err := CastValue(a, "-")
	if err != nil {
		return nil, err
	}

	return aVal.Neg()
}

func Add(a, b interface{}) (*Value, error) {
	aVal, err := CastValue(a, "+")
	if err != nil {
		return nil, err
	}

	bVal, err := CastValue(b, "+")
	if err != nil {
		return nil, err
	}

	return aVal.Add(bVal)
}

func Sub(a, b interface{}) (*Value, error) {
	aVal, err := CastValue(a, "-")
	if err != nil {
		return nil, err
	}

	bVal, err := CastValue(b, "-")
	if err != nil {
		return nil, err
	}

	return aVal.Sub(bVal)
}

func Mul(a, b interface{}) (*Value, error) {
	aVal, err := CastValue(a, "*")
	if err != nil {
		return nil, err
	}

	bVal, err := CastValue(b, "*")
	if err != nil {
		return nil, err
	}

	return aVal.Mul(bVal)
}

func Div(a, b interface{}) (*Value, error) {
	aVal, err := CastValue(a, "/")
	if err != nil {
		return nil, err
	}

	bVal, err := CastValue(b, "/")
	if err != nil {
		return nil, err
	}

	return aVal.Div(bVal)
}
