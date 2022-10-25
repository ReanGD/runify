package calculator_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/ReanGD/runify/server/provider/calculator"
	"github.com/ReanGD/runify/server/system"
)

type CalcSuite struct {
	suite.Suite

	executer *calculator.Executer
}

func (s *CalcSuite) SetupSuite() {
	s.executer = calculator.NewExecuter()
}

func (s *CalcSuite) TearDownSuite() {
	s.executer = nil
}

func (s *CalcSuite) runTest(expression string, expected decimal.Decimal, expectedErrCode system.Error) {
	testName := strings.ReplaceAll(expression, " ", "")
	executer := s.executer
	s.T().Run(testName, func(t *testing.T) {
		res, astErr, err := executer.Execute(expression)
		if expectedErrCode != system.Success {
			require.Error(t, err, expression)
			require.Equal(t, expectedErrCode, astErr.Code(), expression)
		} else {
			require.NoError(t, err, fmt.Errorf("Error in expr: '%s'", expression))
			require.Nil(t, astErr, fmt.Errorf("Ast error in expr: '%s'", expression))
			assertEqualDecimal(t, expected, res, expression)
		}
	})
}

func (s *CalcSuite) runTestsFromArr(data []testDataStr) {
	t := s.T()
	for _, item := range data {
		expression := item.expression
		result, err := decimal.NewFromString(item.result)
		require.NoError(t, err, expression)
		s.runTest(expression, result, system.Success)
	}
}

func (s *CalcSuite) TestGenerated() {
	gen := newTestDataGenerator(time.Now().UnixNano())
	for i := 0; i != 45_000_000; i++ {
		s.runTest(gen.next())
	}
}

func (s *CalcSuite) TestCornerCases() {
	s.runTestsFromArr([]testDataStr{
		{"0123", "123"},
		{"+(0)", "0"},
	})
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
