package calculator_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/ReanGD/runify/server/provider/calculator/ast"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/lexer"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/parser"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CalcSuite struct {
	suite.Suite

	parser *parser.Parser
}

func (s *CalcSuite) SetupSuite() {
	s.parser = parser.NewParser()
}

func (s *CalcSuite) TearDownSuite() {
	s.parser = nil
}

func (s *CalcSuite) runTest(expression string, expected decimal.Decimal) {
	testName := strings.ReplaceAll(expression, " ", "")
	parser := s.parser
	s.T().Run(testName, func(t *testing.T) {
		s := lexer.NewLexer([]byte(expression))

		res, err := parser.Parse(s)
		require.NoError(t, err, fmt.Errorf("Error in expr: '%s'", expression))

		typedRes, ok := res.(*ast.Value)
		require.True(t, ok, expression)

		assertEqualDecimal(t, expected, typedRes.Value(), expression)
	})
}

func (s *CalcSuite) runTestsFromArr(data []testDataStr) {
	t := s.T()
	for _, item := range data {
		expression := item.expression
		result, err := decimal.NewFromString(item.result)
		require.NoError(t, err, expression)
		s.runTest(expression, result)
	}
}

func (s *CalcSuite) TestGenerated() {
	gen := newTestDataGenerator(time.Now().UnixNano())

	for i := 0; i != 1_000_000; i++ {
		result, expression, err := gen.next()
		if err == nil {
			s.runTest(expression, result)
		}
	}
}

func (s *CalcSuite) TestSumAndSubForNegativeAndPositiveInt() {
	s.runTestsFromArr([]testDataStr{
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
