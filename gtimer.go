package timeTick

import (
	"math"
	"time"
)

const (
	STATUS_READY            = 0
	STATUS_RUNNING          = 1
	STATUS_STOPPED          = 2
	STATUS_CLOSED           = -1
	gPANIC_EXIT             = "exit"
	gDEFAULT_TIMES          = math.MaxInt32
	gDEFAULT_SLOT_NUMBER    = 10
	gDEFAULT_WHEEL_INTERVAL = 50
	gDEFAULT_WHEEL_LEVEL    = 6
)

var (
	// 默认的wheel管理对象
	// slots = 10; wheel_interval = 50ms; level = 6
	defaultTimer = New(gDEFAULT_SLOT_NUMBER, gDEFAULT_WHEEL_INTERVAL*time.Millisecond, gDEFAULT_WHEEL_LEVEL)
)

//增加
func Add(interval time.Duration, job JobFunc) *Entry {
	return defaultTimer.Add(interval, job)
}
