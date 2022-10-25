package calculator

import (
	"fmt"

	"github.com/ReanGD/runify/server/provider/calculator/ast"
	goccErr "github.com/ReanGD/runify/server/provider/calculator/gocc/errors"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/lexer"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/parser"
	"github.com/shopspring/decimal"
)

type Executer struct {
	parser *parser.Parser
}

func NewExecuter() *Executer {
	return &Executer{
		parser: parser.NewParser(),
	}
}

func (e *Executer) Execute(expression string) (decimal.Decimal, *ast.Error, error) {
	lexer := lexer.NewLexer([]byte(expression))
	res, err := e.parser.Parse(lexer)
	if err != nil {
		var astError *ast.Error
		parserErr, ok := err.(*goccErr.Error)
		if ok {
			if innerErr, ok := parserErr.Err.(*ast.Error); ok {
				astError = innerErr
			}
		}

		return decimal.Zero, astError, err
	}

	typedRes, ok := res.(*ast.Value)
	if !ok {
		return decimal.Zero, nil, fmt.Errorf("invalid result type; expected *ast.Value, got %T", res)
	}

	return typedRes.Value(), nil, nil
}
