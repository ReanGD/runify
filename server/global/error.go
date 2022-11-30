package global

import "go.uber.org/zap"

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

var errorMap = map[Error]string{
	Success:                  "Success",
	HotkeyParseFailed:        "Hotkey parse failed",
	HotkeyUsesByRunify:       "Hotkey uses by Runify",
	HotkeyUsesByExternalApp:  "Hotkey uses by external app",
	HotkeyBindError:          "Hotkey bind error",
	CalculatorTypeMismatch:   "Calculator type mismatch",
	CalculatorResultTooBig:   "Calculator result too big",
	CalculatorResultTooSmall: "Calculator result too small",
	CalculatorResultRounded:  "Calculator result rounded",
	CalculatorDivisionByZero: "Calculator division by zero",
}

func (e Error) String() string {
	if res, ok := errorMap[e]; ok {
		return res
	}

	return "Unknown error"
}

func (e Error) ZapField() zap.Field {
	return zap.String("ErrorCode", e.String())
}

func (e Error) ZapFieldPrefix(prefix string) zap.Field {
	return zap.String(prefix+"ErrorCode", e.String())
}
