// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(3), // +
			shift(5), // -
			nil,      // *
			nil,      // /
			shift(8), // (
			nil,      // )
			shift(9), // int64
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          // INVALID
			accept(true), // $
			nil,          // +
			nil,          // -
			nil,          // *
			nil,          // /
			nil,          // (
			nil,          // )
			nil,          // int64
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(1), // $, reduce: Calc
			shift(10), // +
			shift(11), // -
			nil,       // *
			nil,       // /
			nil,       // (
			nil,       // )
			nil,       // int64
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(3), // +
			shift(5), // -
			nil,      // *
			nil,      // /
			shift(8), // (
			nil,      // )
			shift(9), // int64
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(4), // $, reduce: Expr0Lvl
			reduce(4), // +, reduce: Expr0Lvl
			reduce(4), // -, reduce: Expr0Lvl
			shift(13), // *
			shift(14), // /
			nil,       // (
			nil,       // )
			nil,       // int64
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(3), // +
			shift(5), // -
			nil,      // *
			nil,      // /
			shift(8), // (
			nil,      // )
			shift(9), // int64
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(7), // $, reduce: Expr1Lvl
			reduce(7), // +, reduce: Expr1Lvl
			reduce(7), // -, reduce: Expr1Lvl
			reduce(7), // *, reduce: Expr1Lvl
			reduce(7), // /, reduce: Expr1Lvl
			nil,       // (
			nil,       // )
			nil,       // int64
		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			reduce(10), // $, reduce: Expr2Lvl
			reduce(10), // +, reduce: Expr2Lvl
			reduce(10), // -, reduce: Expr2Lvl
			reduce(10), // *, reduce: Expr2Lvl
			reduce(10), // /, reduce: Expr2Lvl
			nil,        // (
			nil,        // )
			nil,        // int64
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(17), // +
			shift(19), // -
			nil,       // *
			nil,       // /
			shift(22), // (
			nil,       // )
			shift(23), // int64
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			reduce(12), // $, reduce: Expr3Lvl
			reduce(12), // +, reduce: Expr3Lvl
			reduce(12), // -, reduce: Expr3Lvl
			reduce(12), // *, reduce: Expr3Lvl
			reduce(12), // /, reduce: Expr3Lvl
			nil,        // (
			nil,        // )
			nil,        // int64
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(3), // +
			shift(5), // -
			nil,      // *
			nil,      // /
			shift(8), // (
			nil,      // )
			shift(9), // int64
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(3), // +
			shift(5), // -
			nil,      // *
			nil,      // /
			shift(8), // (
			nil,      // )
			shift(9), // int64
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(9), // $, reduce: Expr2Lvl
			reduce(9), // +, reduce: Expr2Lvl
			reduce(9), // -, reduce: Expr2Lvl
			reduce(9), // *, reduce: Expr2Lvl
			reduce(9), // /, reduce: Expr2Lvl
			nil,       // (
			nil,       // )
			nil,       // int64
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(3), // +
			shift(5), // -
			nil,      // *
			nil,      // /
			shift(8), // (
			nil,      // )
			shift(9), // int64
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,      // INVALID
			nil,      // $
			shift(3), // +
			shift(5), // -
			nil,      // *
			nil,      // /
			shift(8), // (
			nil,      // )
			shift(9), // int64
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(8), // $, reduce: Expr2Lvl
			reduce(8), // +, reduce: Expr2Lvl
			reduce(8), // -, reduce: Expr2Lvl
			reduce(8), // *, reduce: Expr2Lvl
			reduce(8), // /, reduce: Expr2Lvl
			nil,       // (
			nil,       // )
			nil,       // int64
		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(28), // +
			shift(29), // -
			nil,       // *
			nil,       // /
			nil,       // (
			shift(30), // )
			nil,       // int64
		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(17), // +
			shift(19), // -
			nil,       // *
			nil,       // /
			shift(22), // (
			nil,       // )
			shift(23), // int64
		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(4), // +, reduce: Expr0Lvl
			reduce(4), // -, reduce: Expr0Lvl
			shift(32), // *
			shift(33), // /
			nil,       // (
			reduce(4), // ), reduce: Expr0Lvl
			nil,       // int64
		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(17), // +
			shift(19), // -
			nil,       // *
			nil,       // /
			shift(22), // (
			nil,       // )
			shift(23), // int64
		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(7), // +, reduce: Expr1Lvl
			reduce(7), // -, reduce: Expr1Lvl
			reduce(7), // *, reduce: Expr1Lvl
			reduce(7), // /, reduce: Expr1Lvl
			nil,       // (
			reduce(7), // ), reduce: Expr1Lvl
			nil,       // int64
		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			reduce(10), // +, reduce: Expr2Lvl
			reduce(10), // -, reduce: Expr2Lvl
			reduce(10), // *, reduce: Expr2Lvl
			reduce(10), // /, reduce: Expr2Lvl
			nil,        // (
			reduce(10), // ), reduce: Expr2Lvl
			nil,        // int64
		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(17), // +
			shift(19), // -
			nil,       // *
			nil,       // /
			shift(22), // (
			nil,       // )
			shift(23), // int64
		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			reduce(12), // +, reduce: Expr3Lvl
			reduce(12), // -, reduce: Expr3Lvl
			reduce(12), // *, reduce: Expr3Lvl
			reduce(12), // /, reduce: Expr3Lvl
			nil,        // (
			reduce(12), // ), reduce: Expr3Lvl
			nil,        // int64
		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(2), // $, reduce: Expr0Lvl
			reduce(2), // +, reduce: Expr0Lvl
			reduce(2), // -, reduce: Expr0Lvl
			shift(13), // *
			shift(14), // /
			nil,       // (
			nil,       // )
			nil,       // int64
		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(3), // $, reduce: Expr0Lvl
			reduce(3), // +, reduce: Expr0Lvl
			reduce(3), // -, reduce: Expr0Lvl
			shift(13), // *
			shift(14), // /
			nil,       // (
			nil,       // )
			nil,       // int64
		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(5), // $, reduce: Expr1Lvl
			reduce(5), // +, reduce: Expr1Lvl
			reduce(5), // -, reduce: Expr1Lvl
			reduce(5), // *, reduce: Expr1Lvl
			reduce(5), // /, reduce: Expr1Lvl
			nil,       // (
			nil,       // )
			nil,       // int64
		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(6), // $, reduce: Expr1Lvl
			reduce(6), // +, reduce: Expr1Lvl
			reduce(6), // -, reduce: Expr1Lvl
			reduce(6), // *, reduce: Expr1Lvl
			reduce(6), // /, reduce: Expr1Lvl
			nil,       // (
			nil,       // )
			nil,       // int64
		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(17), // +
			shift(19), // -
			nil,       // *
			nil,       // /
			shift(22), // (
			nil,       // )
			shift(23), // int64
		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(17), // +
			shift(19), // -
			nil,       // *
			nil,       // /
			shift(22), // (
			nil,       // )
			shift(23), // int64
		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			reduce(11), // $, reduce: Expr3Lvl
			reduce(11), // +, reduce: Expr3Lvl
			reduce(11), // -, reduce: Expr3Lvl
			reduce(11), // *, reduce: Expr3Lvl
			reduce(11), // /, reduce: Expr3Lvl
			nil,        // (
			nil,        // )
			nil,        // int64
		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(9), // +, reduce: Expr2Lvl
			reduce(9), // -, reduce: Expr2Lvl
			reduce(9), // *, reduce: Expr2Lvl
			reduce(9), // /, reduce: Expr2Lvl
			nil,       // (
			reduce(9), // ), reduce: Expr2Lvl
			nil,       // int64
		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(17), // +
			shift(19), // -
			nil,       // *
			nil,       // /
			shift(22), // (
			nil,       // )
			shift(23), // int64
		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(17), // +
			shift(19), // -
			nil,       // *
			nil,       // /
			shift(22), // (
			nil,       // )
			shift(23), // int64
		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(8), // +, reduce: Expr2Lvl
			reduce(8), // -, reduce: Expr2Lvl
			reduce(8), // *, reduce: Expr2Lvl
			reduce(8), // /, reduce: Expr2Lvl
			nil,       // (
			reduce(8), // ), reduce: Expr2Lvl
			nil,       // int64
		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(28), // +
			shift(29), // -
			nil,       // *
			nil,       // /
			nil,       // (
			shift(40), // )
			nil,       // int64
		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(2), // +, reduce: Expr0Lvl
			reduce(2), // -, reduce: Expr0Lvl
			shift(32), // *
			shift(33), // /
			nil,       // (
			reduce(2), // ), reduce: Expr0Lvl
			nil,       // int64
		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(3), // +, reduce: Expr0Lvl
			reduce(3), // -, reduce: Expr0Lvl
			shift(32), // *
			shift(33), // /
			nil,       // (
			reduce(3), // ), reduce: Expr0Lvl
			nil,       // int64
		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(5), // +, reduce: Expr1Lvl
			reduce(5), // -, reduce: Expr1Lvl
			reduce(5), // *, reduce: Expr1Lvl
			reduce(5), // /, reduce: Expr1Lvl
			nil,       // (
			reduce(5), // ), reduce: Expr1Lvl
			nil,       // int64
		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(6), // +, reduce: Expr1Lvl
			reduce(6), // -, reduce: Expr1Lvl
			reduce(6), // *, reduce: Expr1Lvl
			reduce(6), // /, reduce: Expr1Lvl
			nil,       // (
			reduce(6), // ), reduce: Expr1Lvl
			nil,       // int64
		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			reduce(11), // +, reduce: Expr3Lvl
			reduce(11), // -, reduce: Expr3Lvl
			reduce(11), // *, reduce: Expr3Lvl
			reduce(11), // /, reduce: Expr3Lvl
			nil,        // (
			reduce(11), // ), reduce: Expr3Lvl
			nil,        // int64
		},
	},
}
