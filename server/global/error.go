package global

type Error uint32

const (
	Success                  Error = 0
	HotkeyParseFailed        Error = 1
	HotkeyUsesByRunify       Error = 2
	HotkeyUsesByExternalApp  Error = 3
	HotkeyBindError          Error = 4
	CalculatorTypeMismatch   Error = 5
	CalculatorResultTooBig   Error = 6
	CalculatorResultTooSmall Error = 7
	CalculatorResultRounded  Error = 8
	CalculatorDivisionByZero Error = 9
)
