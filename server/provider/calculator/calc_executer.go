package calculator

import (
	"fmt"

	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/provider/calculator/ast"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/lexer"
	"github.com/ReanGD/runify/server/provider/calculator/gocc/parser"
	"github.com/cockroachdb/apd/v3"
)

type Result struct {
	Value       apd.Decimal
	ParserErr   error
	CalcErrCode global.Error
	Condition   apd.Condition
}

func newResult(parserErr error, astCtx *ast.AstContext) *Result {
	return &Result{
		Value:     apd.Decimal{},
		ParserErr: parserErr,
		Condition: 0,
	}
}

func (r *Result) IsExprValid() bool {
	return r.ParserErr == nil
}

func (r *Result) IsValueValid() bool {
	return r.ParserErr == nil && r.Condition == 0
}

func (r *Result) UserResult() string {
	if r.ParserErr != nil {
		return ""
	}

	if r.Condition == 0 {
		return r.Value.String()
	}

	return r.CalcErrCode.String()
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

func (e *Executer) Execute(expression string) *Result {
	e.ctx.Reset()
	lexer := lexer.NewLexer([]byte(expression))
	parserRes, parserErr := e.parser.Parse(lexer)
	condition, calcErrCode := e.ctx.Error()
	var value apd.Decimal
	if parserErr == nil && condition == 0 {
		if astValue, ok := parserRes.(*ast.Value); ok {
			value = astValue.Value()
		} else {
			parserErr = fmt.Errorf("invalid result type; expected *ast.Value, got %T", parserRes)
		}
	}

	return &Result{
		Value:       value,
		ParserErr:   parserErr,
		CalcErrCode: calcErrCode,
		Condition:   condition,
	}
}
