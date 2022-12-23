package interpreter_test

import (
	"strings"
	"testing"

	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/interpreter"
	"github.com/cockroachdb/apd/v3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CalcSuite struct {
	interpreter *interpreter.Interpreter

	suite.Suite
}

func (s *CalcSuite) SetupSuite() {
	s.interpreter = interpreter.New()
}

func (s *CalcSuite) TearDownSuite() {
	s.interpreter = nil
}

func (s *CalcSuite) runTest(expression string, expectedValue apd.Decimal, expectedCondition apd.Condition) {
	testName := strings.ReplaceAll(expression, " ", "")
	interp := s.interpreter
	s.T().Run(testName, func(t *testing.T) {
		actualRes := interp.Execute(expression)
		require.Equal(t, expectedCondition, actualRes.Condition, expression)
		require.True(t, actualRes.IsExprValid())
		require.NoError(t, actualRes.ParserErr, expression)
		if expectedCondition != 0 {
			require.False(t, actualRes.IsValueValid(), expression)
			require.NotEqual(t, global.Success, actualRes.CalcErrCode, expression)
		} else {
			require.True(t, actualRes.IsValueValid(), expression)
			require.Equal(t, global.Success, actualRes.CalcErrCode, expression)
			assertEqualDecimal(t, expectedValue, actualRes.Value, expression)
		}
	})
}

func (s *CalcSuite) runTestsFromArr(data []testDataStr) {
	t := s.T()
	dctx := s.interpreter.GetApdContext()
	for _, item := range data {
		expression := item.expression
		expectedValue, _, err := dctx.NewFromString(item.result)
		require.NoError(t, err, expression)
		s.runTest(expression, *expectedValue, 0)
	}
}

func (s *CalcSuite) TestGenerated() {
	timer := global.NewTimer()
	gen := newTestDataGenerator(int64(timer), s.interpreter.GetApdContext())
	for i := 0; i != 30_000; i++ {
		s.runTest(gen.next())
	}
	gen.showStats(timer.Duration())
}

func (s *CalcSuite) TestParse() {
	s.runTestsFromArr([]testDataStr{
		{"0123", "123"},
		{"0123.45", "123.45"},
		{"0123.", "123"},
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
		{"1 + 2**2*2", "9"}, // 1 + (2 ** 2) * 2
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
		{"1.1 + 2.2 - 3.3", "0"},
		{"1.1 + 2.2 * 3.3", "8.36"},
		{"1.1 + 6 / 3.3", "2.9181818181818184"},
		{"1.1 + -2.2", "-1.1"},
		{"1.1 + 2.2**2.2", "6.766695778750079"},
		{"6.1 * 2.2 / 3.3", "4.066666666666667"},
		{"3.1 * -2.2", "-6.82"},
		{"3.1 * 2.2**2.3", "19.007887798933407"},
		{"(1.1 + 2.2) * 3.3", "10.89"},
		{"(1.1 + 5.5) / 3", "2.2"},
		{"(1.1 + 2.2)**2", "10.89"},
		{"-(1.1 + 2.2)**2", "-10.89"}, // (-(1.1 + 2.2))**2
	})
}

func (s *CalcSuite) TestError() {
	res := s.interpreter.Execute("1 + qwe")
	s.False(res.IsExprValid())
	s.False(res.IsValueValid())

	res = s.interpreter.Execute("1 +")
	s.False(res.IsExprValid())
	s.False(res.IsValueValid())

	res = s.interpreter.Execute("1.2.3")
	s.False(res.IsExprValid())
	s.False(res.IsValueValid())

	res = s.interpreter.Execute("1/0q")
	s.False(res.IsExprValid())
	s.False(res.IsValueValid())

	res = s.interpreter.Execute("1/0")
	s.True(res.IsExprValid())
	s.False(res.IsValueValid())
}

func TestCalcSuite(t *testing.T) {
	suite.Run(t, new(CalcSuite))
}
