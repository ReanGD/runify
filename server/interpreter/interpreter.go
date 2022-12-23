package interpreter

import (
	"fmt"

	"github.com/ReanGD/runify/server/global"
	"github.com/ReanGD/runify/server/interpreter/ast"
	"github.com/ReanGD/runify/server/interpreter/gocc/lexer"
	"github.com/ReanGD/runify/server/interpreter/gocc/parser"
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

type Interpreter struct {
	ctx    *ast.AstContext
	parser *parser.Parser
}

func New() *Interpreter {
	ctx := ast.NewDefaultAstContext()
	parser := parser.NewParser()
	parser.Context = ctx
	return &Interpreter{
		parser: parser,
		ctx:    ctx,
	}
}

func NewWithCtx(ctx *ast.AstContext) *Interpreter {
	parser := parser.NewParser()
	parser.Context = ctx
	return &Interpreter{
		parser: parser,
		ctx:    ctx,
	}
}

func (i *Interpreter) GetApdContext() apd.Context {
	return i.ctx.GetApdContext()
}

func (i *Interpreter) Execute(expression string) *Result {
	i.ctx.Reset()
	lexer := lexer.NewLexer([]byte(expression))
	parserRes, parserErr := i.parser.Parse(lexer)
	condition, calcErrCode := i.ctx.Error()
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
