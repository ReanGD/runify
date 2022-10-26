package calculator_test

import (
	"fmt"
	"strings"
	"testing"

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
	timer := system.NewTimer()
	gen := newTestDataGenerator(int64(timer))
	for i := 0; i != 30_000; i++ {
		s.runTest(gen.next())
	}
	gen.showStats(timer.Duration())
}

func (s *CalcSuite) TestParse() {
	s.runTestsFromArr([]testDataStr{
		{"0123", "123"},
		{"-+1", "-1"},
		{"+-1", "-1"},
		{"-+-1", "1"},
		{"+-+1", "-1"},
		{"-1-+-+1", "0"},
	})
}

func (s *CalcSuite) TestPriority() {
	s.runTestsFromArr([]testDataStr{
		{"1 + 2 - 3", "0"},
		{"1 + 2 * 3", "7"},
		{"1 + 6 / 3", "3"},
		{"1 + -2", "-1"},
		{"1 + 2**2", "5"},
		{"6 * 2 / 3", "4"},
		{"3 * -2", "-6"},
		{"3 * 2**2", "12"},
		{"-2**2", "-4"},       // -(2 ^ 2)
		{"2 ^ 3 ^ 4", "4096"}, // (2 ^ 3) ^ 4
		{"(1 + 2) * 3", "9"},
		{"(1 + 5) / 3", "2"},
		{"(1 + 2)**2", "9"},
		{"-(1 + 2)**2", "-9"}, // (-(1 + 2))**2
		{"2 ^ (3 ^ 4)", "2417851639229258349412352"},
	})
}

func TestCalcSuite(t *testing.T) {
	suite.Run(t, new(CalcSuite))
}
