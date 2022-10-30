package system

type Error uint32

const (
	Success                   Error = 0
	ShortcutParseFailed       Error = 1
	ShortcutUsesByRunify      Error = 2
	ShortcutUsesByExternalApp Error = 3
	ShortcutBindError         Error = 4
	CalculatorTypeMismatch    Error = 5
	CalculatorResultTooBig    Error = 6
	CalculatorResultTooSmall  Error = 7
	CalculatorResultRounded   Error = 8
	CalculatorDivisionByZero  Error = 9
)
