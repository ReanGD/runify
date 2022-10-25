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
			nil,       // INVALID
			nil,       // $
			shift(3),  // +
			shift(5),  // -
			nil,       // *
			nil,       // /
			shift(8),  // int
			shift(10), // (
			nil,       // )
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
			nil,          // int
			nil,          // (
			nil,          // )
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(1), // $, reduce: Calc
			shift(11), // +
			shift(12), // -
			nil,       // *
			nil,       // /
			nil,       // int
			nil,       // (
			nil,       // )
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(3),  // +
			shift(5),  // -
			nil,       // *
			nil,       // /
			shift(8),  // int
			shift(10), // (
			nil,       // )
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(4), // $, reduce: Expr0Lvl
			reduce(4), // +, reduce: Expr0Lvl
			reduce(4), // -, reduce: Expr0Lvl
			shift(14), // *
			shift(15), // /
			nil,       // int
			nil,       // (
			nil,       // )
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(3),  // +
			shift(5),  // -
			nil,       // *
			nil,       // /
			shift(8),  // int
			shift(10), // (
			nil,       // )
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
			nil,       // int
			nil,       // (
			nil,       // )
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
			nil,        // int
			nil,        // (
			nil,        // )
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			reduce(11), // $, reduce: Number
			reduce(11), // +, reduce: Number
			reduce(11), // -, reduce: Number
			reduce(11), // *, reduce: Number
			reduce(11), // /, reduce: Number
			nil,        // int
			nil,        // (
			nil,        // )
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			reduce(12), // $, reduce: Number
			reduce(12), // +, reduce: Number
			reduce(12), // -, reduce: Number
			reduce(12), // *, reduce: Number
			reduce(12), // /, reduce: Number
			nil,        // int
			nil,        // (
			nil,        // )
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(18), // +
			shift(20), // -
			nil,       // *
			nil,       // /
			shift(23), // int
			shift(25), // (
			nil,       // )
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(3),  // +
			shift(5),  // -
			nil,       // *
			nil,       // /
			shift(8),  // int
			shift(10), // (
			nil,       // )
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(3),  // +
			shift(5),  // -
			nil,       // *
			nil,       // /
			shift(8),  // int
			shift(10), // (
			nil,       // )
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(9), // $, reduce: Expr2Lvl
			reduce(9), // +, reduce: Expr2Lvl
			reduce(9), // -, reduce: Expr2Lvl
			reduce(9), // *, reduce: Expr2Lvl
			reduce(9), // /, reduce: Expr2Lvl
			nil,       // int
			nil,       // (
			nil,       // )
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(3),  // +
			shift(5),  // -
			nil,       // *
			nil,       // /
			shift(8),  // int
			shift(10), // (
			nil,       // )
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(3),  // +
			shift(5),  // -
			nil,       // *
			nil,       // /
			shift(8),  // int
			shift(10), // (
			nil,       // )
		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(8), // $, reduce: Expr2Lvl
			reduce(8), // +, reduce: Expr2Lvl
			reduce(8), // -, reduce: Expr2Lvl
			reduce(8), // *, reduce: Expr2Lvl
			reduce(8), // /, reduce: Expr2Lvl
			nil,       // int
			nil,       // (
			nil,       // )
		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(30), // +
			shift(31), // -
			nil,       // *
			nil,       // /
			nil,       // int
			nil,       // (
			shift(32), // )
		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(18), // +
			shift(20), // -
			nil,       // *
			nil,       // /
			shift(23), // int
			shift(25), // (
			nil,       // )
		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(4), // +, reduce: Expr0Lvl
			reduce(4), // -, reduce: Expr0Lvl
			shift(34), // *
			shift(35), // /
			nil,       // int
			nil,       // (
			reduce(4), // ), reduce: Expr0Lvl
		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(18), // +
			shift(20), // -
			nil,       // *
			nil,       // /
			shift(23), // int
			shift(25), // (
			nil,       // )
		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(7), // +, reduce: Expr1Lvl
			reduce(7), // -, reduce: Expr1Lvl
			reduce(7), // *, reduce: Expr1Lvl
			reduce(7), // /, reduce: Expr1Lvl
			nil,       // int
			nil,       // (
			reduce(7), // ), reduce: Expr1Lvl
		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			reduce(10), // +, reduce: Expr2Lvl
			reduce(10), // -, reduce: Expr2Lvl
			reduce(10), // *, reduce: Expr2Lvl
			reduce(10), // /, reduce: Expr2Lvl
			nil,        // int
			nil,        // (
			reduce(10), // ), reduce: Expr2Lvl
		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			reduce(11), // +, reduce: Number
			reduce(11), // -, reduce: Number
			reduce(11), // *, reduce: Number
			reduce(11), // /, reduce: Number
			nil,        // int
			nil,        // (
			reduce(11), // ), reduce: Number
		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			reduce(12), // +, reduce: Number
			reduce(12), // -, reduce: Number
			reduce(12), // *, reduce: Number
			reduce(12), // /, reduce: Number
			nil,        // int
			nil,        // (
			reduce(12), // ), reduce: Number
		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(18), // +
			shift(20), // -
			nil,       // *
			nil,       // /
			shift(23), // int
			shift(25), // (
			nil,       // )
		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(2), // $, reduce: Expr0Lvl
			reduce(2), // +, reduce: Expr0Lvl
			reduce(2), // -, reduce: Expr0Lvl
			shift(14), // *
			shift(15), // /
			nil,       // int
			nil,       // (
			nil,       // )
		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(3), // $, reduce: Expr0Lvl
			reduce(3), // +, reduce: Expr0Lvl
			reduce(3), // -, reduce: Expr0Lvl
			shift(14), // *
			shift(15), // /
			nil,       // int
			nil,       // (
			nil,       // )
		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(5), // $, reduce: Expr1Lvl
			reduce(5), // +, reduce: Expr1Lvl
			reduce(5), // -, reduce: Expr1Lvl
			reduce(5), // *, reduce: Expr1Lvl
			reduce(5), // /, reduce: Expr1Lvl
			nil,       // int
			nil,       // (
			nil,       // )
		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(6), // $, reduce: Expr1Lvl
			reduce(6), // +, reduce: Expr1Lvl
			reduce(6), // -, reduce: Expr1Lvl
			reduce(6), // *, reduce: Expr1Lvl
			reduce(6), // /, reduce: Expr1Lvl
			nil,       // int
			nil,       // (
			nil,       // )
		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(18), // +
			shift(20), // -
			nil,       // *
			nil,       // /
			shift(23), // int
			shift(25), // (
			nil,       // )
		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(18), // +
			shift(20), // -
			nil,       // *
			nil,       // /
			shift(23), // int
			shift(25), // (
			nil,       // )
		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			reduce(13), // $, reduce: Bracket
			reduce(13), // +, reduce: Bracket
			reduce(13), // -, reduce: Bracket
			reduce(13), // *, reduce: Bracket
			reduce(13), // /, reduce: Bracket
			nil,        // int
			nil,        // (
			nil,        // )
		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(9), // +, reduce: Expr2Lvl
			reduce(9), // -, reduce: Expr2Lvl
			reduce(9), // *, reduce: Expr2Lvl
			reduce(9), // /, reduce: Expr2Lvl
			nil,       // int
			nil,       // (
			reduce(9), // ), reduce: Expr2Lvl
		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(18), // +
			shift(20), // -
			nil,       // *
			nil,       // /
			shift(23), // int
			shift(25), // (
			nil,       // )
		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(18), // +
			shift(20), // -
			nil,       // *
			nil,       // /
			shift(23), // int
			shift(25), // (
			nil,       // )
		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(8), // +, reduce: Expr2Lvl
			reduce(8), // -, reduce: Expr2Lvl
			reduce(8), // *, reduce: Expr2Lvl
			reduce(8), // /, reduce: Expr2Lvl
			nil,       // int
			nil,       // (
			reduce(8), // ), reduce: Expr2Lvl
		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(30), // +
			shift(31), // -
			nil,       // *
			nil,       // /
			nil,       // int
			nil,       // (
			shift(42), // )
		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(2), // +, reduce: Expr0Lvl
			reduce(2), // -, reduce: Expr0Lvl
			shift(34), // *
			shift(35), // /
			nil,       // int
			nil,       // (
			reduce(2), // ), reduce: Expr0Lvl
		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(3), // +, reduce: Expr0Lvl
			reduce(3), // -, reduce: Expr0Lvl
			shift(34), // *
			shift(35), // /
			nil,       // int
			nil,       // (
			reduce(3), // ), reduce: Expr0Lvl
		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(5), // +, reduce: Expr1Lvl
			reduce(5), // -, reduce: Expr1Lvl
			reduce(5), // *, reduce: Expr1Lvl
			reduce(5), // /, reduce: Expr1Lvl
			nil,       // int
			nil,       // (
			reduce(5), // ), reduce: Expr1Lvl
		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(6), // +, reduce: Expr1Lvl
			reduce(6), // -, reduce: Expr1Lvl
			reduce(6), // *, reduce: Expr1Lvl
			reduce(6), // /, reduce: Expr1Lvl
			nil,       // int
			nil,       // (
			reduce(6), // ), reduce: Expr1Lvl
		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			reduce(13), // +, reduce: Bracket
			reduce(13), // -, reduce: Bracket
			reduce(13), // *, reduce: Bracket
			reduce(13), // /, reduce: Bracket
			nil,        // int
			nil,        // (
			reduce(13), // ), reduce: Bracket
		},
	},
}
