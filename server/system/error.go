package system

type Error uint32

const (
	Success                   Error = 0
	ShortcutParseFailed       Error = 1
	ShortcutUsesByRunify      Error = 2
	ShortcutUsesByExternalApp Error = 3
	ShortcutBindError         Error = 4
	CalculatorDivisionByZero  Error = 5
)
