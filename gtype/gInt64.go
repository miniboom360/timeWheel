package gtype

import "sync/atomic"

type Int64 struct {
	val int64
}

func NewInt64(value ...int64) *Int64 {
	if len(value) > 0 {
		return &Int64{val: value[0]}
	}
	return &Int64{}
}

func (t *Int64) Clone() *Int64 {
	return NewInt64(t.Val())
}

func (t *Int64) Set(value int64) (old int64) {
	return atomic.SwapInt64(&t.val, value)
}

func (t *Int64) Val() int64 {
	return atomic.LoadInt64(&t.val)
}

func (t *Int64) Add(delta int64) int64 {
	return atomic.AddInt64(&t.val, delta)
}
