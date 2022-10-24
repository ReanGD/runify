package calculator_test

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func assertEqualDecimal(t *testing.T, expected decimal.Decimal, actual decimal.Decimal, msgAndArgs ...interface{}) {
	t.Helper()
	if !expected.Equal(actual) {
		assert.Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected.String(), actual.String()), msgAndArgs...)
		t.FailNow()
	}
}

const (
	maxChance        = 1000
	firstAction      = 600
	secondAction     = 200
	secondActionRare = 40
)

type testDataStr struct {
	expression string
	result     string
}

type testDataGenerator struct {
	gen *rand.Rand
}

func newTestDataGenerator(seed int64) *testDataGenerator {
	return &testDataGenerator{
		gen: rand.New(rand.NewSource(seed)),
	}
}

func (s *testDataGenerator) next() (result decimal.Decimal, expression string, err error) {
	defer func() {
		if recoverResult := recover(); recoverResult != nil {
			switch recoverVal := recoverResult.(type) {
			case string:
				if strings.Contains(recoverVal, "division by 0") {
					err = errors.New("division by 0")
				} else if strings.Contains(recoverVal, "overflow") {
					err = errors.New("overflow")
				} else {
					err = errors.New("unknown error")
				}
			case error:
				err = errors.New("unknown error")
			default:
				err = errors.New("unknown error")
			}
		} else {
			err = errors.New("unknown error")
		}
	}()

	result, expression = s.genExprAddSub()
	return
}

func (s *testDataGenerator) genExprAddSub() (decimal.Decimal, string) {
	kind := s.gen.Intn(maxChance)
	if kind < firstAction {
		return s.genExprMulDiv()
	}

	left, leftStr := s.genExprMulDiv()
	right, rightStr := s.genExprMulDiv()
	kind -= firstAction
	if kind < secondAction {
		res := left.Add(right)
		return res, fmt.Sprintf("%s + %s", leftStr, rightStr)
	}

	res := left.Sub(right)
	return res, fmt.Sprintf("%s - %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprMulDiv() (decimal.Decimal, string) {
	kind := s.gen.Intn(maxChance)
	if kind < firstAction {
		return s.genExprUnaryPlusMinus()
	}

	left, leftStr := s.genExprUnaryPlusMinus()
	right, rightStr := s.genExprUnaryPlusMinus()
	kind -= firstAction
	if kind < secondAction {
		res := left.Mul(right)
		return res, fmt.Sprintf("%s * %s", leftStr, rightStr)
	}

	res := left.Div(right)
	return res, fmt.Sprintf("%s / %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprUnaryPlusMinus() (decimal.Decimal, string) {
	kind := s.gen.Intn(maxChance)
	if kind < firstAction {
		return s.genExprBracketsInt()
	}

	res, resStr := s.genExprBracketsInt()
	kind -= firstAction
	if kind < secondActionRare {
		return res.Neg(), fmt.Sprintf("-%s", resStr)
	}

	return res, fmt.Sprintf("+%s", resStr)
}

func (s *testDataGenerator) genExprBracketsInt() (decimal.Decimal, string) {
	kind := s.gen.Intn(maxChance)
	if kind < firstAction {
		return s.genInt64()
	}

	res, resStr := s.genExprAddSub()
	return res, fmt.Sprintf("(%s)", resStr)
}

func (s *testDataGenerator) genInt64() (decimal.Decimal, string) {
	res := decimal.NewFromInt32(s.gen.Int31())
	return res, res.String()
}
