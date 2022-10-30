package calculator

import (
	"fmt"

	"github.com/ReanGD/runify/server/provider/calculator/ast"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/lexer"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/parser"
	"github.com/ReanGD/runify/server/system"
	"github.com/cockroachdb/apd/v3"
)

type Result struct {
	Value         *ast.Value
	ParserErr     error
	SystemErrCode system.Error
	Condition     apd.Condition
}

type Executer struct {
	ctx    *ast.AstContext
	parser *parser.Parser
}

func NewExecuter() *Executer {
	ctx := ast.NewDefaultAstContext()
	parser := parser.NewParser()
	parser.Context = ctx
	return &Executer{
		parser: parser,
		ctx:    ctx,
	}
}

func NewExecuterWithCtx(ctx *ast.AstContext) *Executer {
	parser := parser.NewParser()
	parser.Context = ctx
	return &Executer{
		parser: parser,
		ctx:    ctx,
	}
}

func (e *Executer) GetApdContext() apd.Context {
	return e.ctx.GetApdContext()
}

func (e *Executer) Execute(expression string) Result {
	e.ctx.Reset()
	lexer := lexer.NewLexer([]byte(expression))
	parserRes, parserErr := e.parser.Parse(lexer)

	var res Result
	res.Value = nil
	res.ParserErr = parserErr
	res.Condition, res.SystemErrCode = e.ctx.Error()

	if parserErr == nil {
		var ok bool
		if res.Value, ok = parserRes.(*ast.Value); !ok {
			res.ParserErr = fmt.Errorf("invalid result type; expected *ast.Value, got %T", res)
		}
	}

	return res
}
