package system

import (
	"time"
	_ "unsafe" // for go:linkname
)

//go:linkname runtimeNano runtime.nanotime
func runtimeNano() int64

type Time int64

func Now() Time {
	return Time(runtimeNano())
}

func NewEmptyTime() Time {
	return Time(0)
}

func TimeDiff(start Time, finish Time) time.Duration {
	return time.Duration(int64(finish) - int64(start))
}

func (t Time) Add(value time.Duration) Time {
	return t + Time(value.Nanoseconds())
}

func (t Time) Seconds() float64 {
	return time.Duration(t).Seconds()
}

func (t Time) ToInt64() int64 {
	return int64(t)
}

func (t Time) IsEmpty() bool {
	return int64(t) == 0
}

type Timer int64

func NewTimer() Timer {
	return Timer(runtimeNano())
}

func (t Timer) StartTime() Time {
	return Time(t)
}

func (t Timer) Seconds() float64 {
	return time.Duration(runtimeNano() - int64(t)).Seconds()
}

func (t Timer) Duration() time.Duration {
	return time.Duration(runtimeNano() - int64(t))
}
