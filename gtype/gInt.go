package gtype

import "sync/atomic"

type Int struct {
	val int64
}

func NewInt(value ...int) *Int {
	if len(value) > 0 {
		return &Int{val: int64(value[0])}
	}
	return &Int{}
}

func (t *Int) Clone() *Int {
	return NewInt(t.Val())
}

func (t *Int) Val() int {
	return int(atomic.LoadInt64(&t.val))
}

// 并发安全设置变量值，返回之前的旧值
func (t *Int) Set(value int) (old int) {
	return int(atomic.SwapInt64(&t.val, int64(value)))
}

func (t *Int) Add(delta int) int {
	return int(atomic.AddInt64(&t.val, int64(delta)))
}
