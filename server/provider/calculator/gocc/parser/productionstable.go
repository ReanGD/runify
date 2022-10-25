// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"github.com/ReanGD/runify/server/provider/calculator/ast"
)

type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib, interface{}) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Calc	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Calc : Expr0Lvl	<<  >>`,
		Id:         "Calc",
		NTType:     1,
		Index:      1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr0Lvl : Expr0Lvl "+" Expr1Lvl	<< ast.BinaryExpr(X[0], X[2], X[1]) >>`,
		Id:         "Expr0Lvl",
		NTType:     2,
		Index:      2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.BinaryExpr(X[0], X[2], X[1])
		},
	},
	ProdTabEntry{
		String: `Expr0Lvl : Expr0Lvl "-" Expr1Lvl	<< ast.BinaryExpr(X[0], X[2], X[1]) >>`,
		Id:         "Expr0Lvl",
		NTType:     2,
		Index:      3,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.BinaryExpr(X[0], X[2], X[1])
		},
	},
	ProdTabEntry{
		String: `Expr0Lvl : Expr1Lvl	<<  >>`,
		Id:         "Expr0Lvl",
		NTType:     2,
		Index:      4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr1Lvl : Expr1Lvl "*" Expr2Lvl	<< ast.BinaryExpr(X[0], X[2], X[1]) >>`,
		Id:         "Expr1Lvl",
		NTType:     3,
		Index:      5,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.BinaryExpr(X[0], X[2], X[1])
		},
	},
	ProdTabEntry{
		String: `Expr1Lvl : Expr1Lvl "/" Expr2Lvl	<< ast.BinaryExpr(X[0], X[2], X[1]) >>`,
		Id:         "Expr1Lvl",
		NTType:     3,
		Index:      6,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.BinaryExpr(X[0], X[2], X[1])
		},
	},
	ProdTabEntry{
		String: `Expr1Lvl : Expr2Lvl	<<  >>`,
		Id:         "Expr1Lvl",
		NTType:     3,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expr2Lvl : "-" Expr2Lvl	<< ast.UnaryExpr(X[1], X[0]) >>`,
		Id:         "Expr2Lvl",
		NTType:     4,
		Index:      8,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.UnaryExpr(X[1], X[0])
		},
	},
	ProdTabEntry{
		String: `Expr2Lvl : "+" Expr2Lvl	<< ast.UnaryExpr(X[1], X[0]) >>`,
		Id:         "Expr2Lvl",
		NTType:     4,
		Index:      9,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.UnaryExpr(X[1], X[0])
		},
	},
	ProdTabEntry{
		String: `Expr2Lvl : Number	<<  >>`,
		Id:         "Expr2Lvl",
		NTType:     4,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Number : int	<< ast.NewValueFromToken(X[0]) >>`,
		Id:         "Number",
		NTType:     5,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return ast.NewValueFromToken(X[0])
		},
	},
	ProdTabEntry{
		String: `Number : Bracket	<<  >>`,
		Id:         "Number",
		NTType:     5,
		Index:      12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bracket : "(" Expr0Lvl ")"	<< X[1], nil >>`,
		Id:         "Bracket",
		NTType:     6,
		Index:      13,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib, C interface{}) (Attrib, error) {
			return X[1], nil
		},
	},
}
