package timeTick

import (
	"time"
	"timeWheel/gtype"
)

type Timer struct {
	status     *gtype.Int //定时器状态
	wheels     []*wheel   //分层时间轮对象
	length     int        //分层层数
	number     int        //每一层Slot Number
	intervalMs int64      //最小时间刻度(毫秒)
}

//单层时间轮
type wheel struct {
	timer      *Timer        //所属定时器
	level      int           //所属分层索引号
	slots      []*gtype.List //所有的循环任务项，按照Slot Number进行分组
	number     int64         //Slot Number=len(slots)
	ticks      *gtype.Int64  //当前时间轮已转动的刻度数量
	totalMs    int64         //整个时间轮的时间长度(毫秒)=number*interval
	createMs   int64
	intervalMs int64
}

func New(slot int, interval time.Duration, level ...int) *Timer { //todo:slots是动态变化的，该怎么处理？
	length := gDEFAULT_WHEEL_LEVEL
	if len(level) > 0 {
		length = level[0]
	}
	t := &Timer{
		status:     gtype.NewInt(STATUS_RUNNING),
		wheels:     make([]*wheel, length),
		length:     length,
		number:     slot,
		intervalMs: interval.Nanoseconds() / 1e6, //间隔ms
	}

	for i := 0; i < length; i++ {
		if i > 0 { //上一个轮盘的总刻度为下一个轮盘的单位刻度
			n := time.Duration(t.wheels[i-1].totalMs) * time.Millisecond
			w := t.newWheel(i, slot, n)
			t.wheels[i] = w
			//todo:最后一个时间轮盘没有addEntry?
			t.wheels[i-1].addEntry(n, w.proceed, false, gDEFAULT_TIMES, STATUS_READY)
		} else {
			t.wheels[i] = t.newWheel(i, slot, interval)
		}
	}

	return t
}

func (t *Timer) newWheel(level int, slot int, interval time.Duration) *wheel {
	w := &wheel{
		timer:      t,
		level:      level,
		slots:      make([]*gtype.List, slot),
		number:     int64(slot),
		ticks:      gtype.NewInt64(),
		totalMs:    int64(slot) * interval.Nanoseconds() / 1e6,
		createMs:   time.Now().UnixNano() / 1e6,
		intervalMs: interval.Nanoseconds() / 1e6,
	}
	for i := int64(0); i < w.number; i++ {
		w.slots[i] = gtype.NewList()
	}
	return w
}

//添加循环任务
func (t *Timer) Add(interval time.Duration, job JobFunc) *Entry {
	return t.doAddEntry(interval, job, false, gDEFAULT_TIMES, STATUS_READY)
}

// 添加定时任务
func (t *Timer) doAddEntry(interval time.Duration, job JobFunc, singleton bool, time int, status int) *Entry {
	return t.wheels[t.getLevelByIntervalMs(interval.Nanoseconds()/1e6)].addEntry(interval, job, singleton, time, status)
}

// 添加定时任务，给定父级Entry,间隔参数为毫秒
func (t *Timer) doAddEntryByParent(interval int64, parent *Entry) *Entry {
	return t.wheels[t.getLevelByIntervalMs(interval)].addEntryByParent(interval, parent)
}

// 根据intervalMs计算添加的分层索引
func (t *Timer) getLevelByIntervalMs(intervalMs int64) int {
	pos, cmp := t.binSearchIndex(intervalMs)
	switch cmp {
	case 0:
		fallthrough
	// intervalMs比最后匹配值小
	case -1:
		i := pos
		for ; i > 0; i-- {
			if intervalMs > t.wheels[i].intervalMs && intervalMs <= t.wheels[i].totalMs {
				return i
			}
		}
		return i
		// intervalMs比最后匹配值大
	case 1:
		i := pos
		for ; i < t.length-1; i++ {
			if intervalMs > t.wheels[i].intervalMs && intervalMs <= t.wheels[i].totalMs {
				return i
			}
		}
		return i
	}
	return 0
}

// 二分查找当前任务可以添加的时间轮对象索引
func (t *Timer) binSearchIndex(n int64) (index, result int) {
	min := 0
	max := t.length - 1
	mid := 0
	cmp := -2
	for min <= max {
		mid = int((min + max) / 2)
		switch {
		case t.wheels[mid].intervalMs == n:
			cmp = 0
		case t.wheels[mid].intervalMs > n:
			cmp = -1
		case t.wheels[mid].intervalMs < n:
			cmp = 1
		}
		switch cmp {
		case -1:
			max = mid - 1
		case 1:
			min = mid + 1
		case 0:
			return mid, cmp
		}
	}
	return mid, cmp
}
