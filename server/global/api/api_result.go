package api

import "github.com/ReanGD/runify/server/global"

type VoidResult interface {
	SetResult()
}

type FuncVoidResult struct {
	action func()
}

func NewFuncVoidResult(action func()) *FuncVoidResult {
	return &FuncVoidResult{
		action: action,
	}
}

func (r *FuncVoidResult) SetResult() {
	r.action()
}

type ChanVoidResult struct {
	ch chan struct{}
}

func NewChanVoidResult() *ChanVoidResult {
	return &ChanVoidResult{
		ch: make(chan struct{}, 1),
	}
}

func (r *ChanVoidResult) SetResult() {
	r.ch <- struct{}{}
}

type BoolResult interface {
	SetResult(value bool)
}

type FuncBoolResult struct {
	action func(result bool)
}

func NewFuncBoolResult(action func(result bool)) *FuncBoolResult {
	return &FuncBoolResult{
		action: action,
	}
}

func (r *FuncBoolResult) SetResult(value bool) {
	r.action(value)
}

type ChanBoolResult struct {
	ch chan bool
}

func NewChanBoolResult() *ChanBoolResult {
	return &ChanBoolResult{
		ch: make(chan bool, 1),
	}
}

func (r *ChanBoolResult) GetChannel() <-chan bool {
	return r.ch
}

func (r *ChanBoolResult) SetResult(value bool) {
	r.ch <- value
}

type ErrorResult interface {
	SetResult(value error)
}

type FuncErrorResult struct {
	action func(result error)
}

func NewFuncErrorResult(action func(result error)) *FuncErrorResult {
	return &FuncErrorResult{
		action: action,
	}
}

func (r *FuncErrorResult) SetResult(value error) {
	r.action(value)
}

type ChanErrorResult struct {
	ch chan error
}

func NewChanErrorResult() *ChanErrorResult {
	return &ChanErrorResult{
		ch: make(chan error, 1),
	}
}

func (r *ChanErrorResult) SetResult(value error) {
	r.ch <- value
}

func (r *ChanErrorResult) GetChannel() <-chan error {
	return r.ch
}

type ErrorCodeResult interface {
	SetResult(value global.Error)
}

type FuncErrorCodeResult struct {
	action func(result global.Error)
}

func NewFuncErrorCodeResult(action func(result global.Error)) *FuncErrorCodeResult {
	return &FuncErrorCodeResult{
		action: action,
	}
}

func (r *FuncErrorCodeResult) SetResult(value global.Error) {
	r.action(value)
}

type ChanErrorCodeResult struct {
	ch chan global.Error
}

func NewChanErrorCodeResult() *ChanErrorCodeResult {
	return &ChanErrorCodeResult{
		ch: make(chan global.Error, 1),
	}
}

func (r *ChanErrorCodeResult) SetResult(value global.Error) {
	r.ch <- value
}

func (r *ChanErrorCodeResult) GetChannel() <-chan global.Error {
	return r.ch
}
