package calculator_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ReanGD/runify/server/provider/calculator/ast"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/lexer"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/parser"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type testData struct {
	expression string
	result     string
}

func equalDecimal(t *testing.T, expected decimal.Decimal, actual decimal.Decimal, msgAndArgs ...interface{}) {
	t.Helper()
	if !expected.Equal(actual) {
		assert.Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected.String(), actual.String()), msgAndArgs...)
		t.FailNow()
	}
}

type CalcSuite struct {
	suite.Suite
}

func (s *CalcSuite) runTest(tests []testData) {
	for _, tdata := range tests {
		expression := tdata.expression
		testName := strings.ReplaceAll(fmt.Sprintf("%s == %s", expression, tdata.result), " ", "")

		p := parser.NewParser()
		s.T().Run(testName, func(t *testing.T) {
			s := lexer.NewLexer([]byte(expression))
			res, err := p.Parse(s)
			require.NoError(t, err, expression)
			typedVal, ok := res.(*ast.Value)
			require.True(t, ok, expression)
			typedRes, err := decimal.NewFromString(tdata.result)
			require.NoError(t, err, expression)
			equalDecimal(t, typedRes, typedVal.Value(), expression)
		})
	}
}

func (s *CalcSuite) TestSumAndSubForNegativeAndPositiveInt() {
	s.runTest([]testData{
		{"18 + 25", "43"},
		{"+18 + -25", "-7"},
		{"-18 + 25", "7"},
		{"-18 + -25", "-43"},
		{"18 - 25", "-7"},
		{"+18 - -25", "43"},
		{"-18 - 25", "-43"},
		{"-18 - -25", "7"},
		{"181 + 325 + 4", "510"},
		{"181 + 325 - 4", "502"},
		{"181 - 325 + 4", "-140"},
		{"181 - 325 - 4", "-148"},
		{"-181 + 325 + 4", "148"},
		{"-181 + 325 - 4", "140"},
		{"-181 - 325 + 4", "-502"},
		{"-181 - 325 - 4", "-510"},
		{"+181 + -325 + 4", "-140"},
		{"181 + -325 - 4", "-148"},
		{"+181 - -325 + 4", "510"},
		{"181 - -325 - 4", "502"},
		{"-181 + -325 + 4", "-502"},
		{"-181 + -325 - 4", "-510"},
		{"-181 - -325 + 4", "148"},
		{"-181 - -325 - 4", "140"},
		{"378 + 154 + 223 + 923", "1678"},
		{"378 + 154 + 223 - 923", "-168"},
		{"-378 + -154 - 223 + 923", "168"},
		{"-378 + -154 - 223 - 923", "-1678"},
		{"-378 - -154 + 223 + 923", "922"},
		{"-378 - -154 + 223 - 923", "-924"},
		{"-378 - -154 - 223 + 923", "476"},
		{"-378 - -154 - 223 - 923", "-1370"},
	})
}

func TestCalcSuite(t *testing.T) {
	suite.Run(t, new(CalcSuite))
}
