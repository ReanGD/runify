package calculator_test

import (
	"fmt"
	"testing"

	"github.com/ReanGD/runify/server/provider/calculator/ast"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/lexer"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/parser"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CalcSuite struct {
	suite.Suite
}

func (s *CalcSuite) TestInt() {
	var tests = []struct {
		expression string
		result     int64
	}{
		{"1 + 2", 3},
		{"2 - 3", -1},
	}

	for _, tdata := range tests {
		expression := tdata.expression
		testName := fmt.Sprintf("%s == %d", expression, tdata.result)
		p := parser.NewParser()
		s.T().Run(testName, func(t *testing.T) {
			s := lexer.NewLexer([]byte(expression))
			res, err := p.Parse(s)
			require.NoError(t, err, expression)
			typedVal, ok := res.(*ast.Value)
			require.True(t, ok, expression)
			require.Equal(t, tdata.result, typedVal.Value(), expression)
		})
	}
}

func TestCalcSuite(t *testing.T) {
	suite.Run(t, new(CalcSuite))
}
