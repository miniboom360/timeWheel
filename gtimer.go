package timeTick

import "math"

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
