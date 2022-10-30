package calculator_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ReanGD/runify/server/provider/calculator"
	"github.com/ReanGD/runify/server/system"
	"github.com/cockroachdb/apd/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CalcSuite struct {
	executer *calculator.Executer

	suite.Suite
}

func (s *CalcSuite) SetupSuite() {
	s.executer = calculator.NewExecuter()
}

func (s *CalcSuite) TearDownSuite() {
	s.executer = nil
}

func (s *CalcSuite) runTest(expression string, expectedValue apd.Decimal, expectedCondition apd.Condition) {
	testName := strings.ReplaceAll(expression, " ", "")
	executer := s.executer
	s.T().Run(testName, func(t *testing.T) {
		actualRes := executer.Execute(expression)
		require.Equal(t, expectedCondition, actualRes.Condition, expression)
		if expectedCondition != 0 {
			require.Error(t, actualRes.ParserErr, expression)
			require.NotEqual(t, system.Success, actualRes.SystemErrCode, expression)
		} else {
			require.NoError(t, actualRes.ParserErr, fmt.Errorf("Error in expr: '%s'", expression))
			require.Equal(t, system.Success, actualRes.SystemErrCode, expression)
			require.NotNil(t, actualRes.Value, expression)
			assertEqualDecimal(t, expectedValue, actualRes.Value.Value(), expression)
		}
	})
}

func (s *CalcSuite) runTestsFromArr(data []testDataStr) {
	t := s.T()
	dctx := s.executer.GetApdContext()
	for _, item := range data {
		expression := item.expression
		expectedValue, _, err := dctx.NewFromString(item.result)
		require.NoError(t, err, expression)
		s.runTest(expression, *expectedValue, 0)
	}
}

func (s *CalcSuite) TestGenerated() {
	timer := system.NewTimer()
	gen := newTestDataGenerator(int64(timer), s.executer.GetApdContext())
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
