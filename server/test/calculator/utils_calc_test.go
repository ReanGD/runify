package calculator_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ReanGD/runify/server/system"
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
	gen     *rand.Rand
	errCode system.Error
}

func newTestDataGenerator(seed int64) *testDataGenerator {
	return &testDataGenerator{
		gen: rand.New(rand.NewSource(seed)),
	}
}

func (s *testDataGenerator) next() (string, decimal.Decimal, system.Error) {
	s.errCode = system.Success
	result, expression := s.genExprAddSub()
	return expression, result, s.errCode
}

func (s *testDataGenerator) genExprAddSub() (decimal.Decimal, string) {
	kind := s.gen.Intn(maxChance)
	if kind < firstAction {
		return s.genExprMulDiv()
	}

	res := decimal.Zero
	left, leftStr := s.genExprMulDiv()
	right, rightStr := s.genExprMulDiv()
	kind -= firstAction
	if kind < secondAction {
		if s.errCode == system.Success {
			res = left.Add(right)
		}
		return res, fmt.Sprintf("%s + %s", leftStr, rightStr)
	}

	if s.errCode == system.Success {
		res = left.Sub(right)
	}
	return res, fmt.Sprintf("%s - %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprMulDiv() (decimal.Decimal, string) {
	kind := s.gen.Intn(maxChance)
	if kind < firstAction {
		return s.genExprUnaryPlusMinus()
	}

	res := decimal.Zero
	left, leftStr := s.genExprUnaryPlusMinus()
	right, rightStr := s.genExprUnaryPlusMinus()
	kind -= firstAction
	if kind < secondAction {
		if s.errCode == system.Success {
			res = left.Mul(right)
		}
		return res, fmt.Sprintf("%s * %s", leftStr, rightStr)
	}

	if right.Sign() == 0 {
		s.errCode = system.CalculatorDivisionByZero
	}
	if s.errCode == system.Success {
		res = left.Div(right)
	}
	return res, fmt.Sprintf("%s / %s", leftStr, rightStr)
}

func (s *testDataGenerator) genExprUnaryPlusMinus() (decimal.Decimal, string) {
	kind := s.gen.Intn(maxChance)
	if kind < firstAction {
		return s.genExprBracketsInt()
	}

	res := decimal.Zero
	right, rightStr := s.genExprBracketsInt()
	kind -= firstAction
	if kind < secondActionRare {
		if s.errCode == system.Success {
			res = right.Neg()
		}
		return res, fmt.Sprintf("-%s", rightStr)
	}

	if s.errCode == system.Success {
		res = right
	}
	return res, fmt.Sprintf("+%s", rightStr)
}

func (s *testDataGenerator) genExprBracketsInt() (decimal.Decimal, string) {
	kind := s.gen.Intn(maxChance)
	if kind < firstAction {
		return s.genInt64()
	}

	res := decimal.Zero
	inner, innerStr := s.genExprAddSub()
	if s.errCode == system.Success {
		res = inner
	}
	return res, fmt.Sprintf("(%s)", innerStr)
}

func (s *testDataGenerator) genInt64() (decimal.Decimal, string) {
	res := decimal.Zero

	val := decimal.NewFromInt32(s.gen.Int31())
	if s.errCode == system.Success {
		res = val
	}
	return res, val.String()
}
