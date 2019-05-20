package timeTick

import (
	"fmt"
	"github.com/gogf/gf/g/os/gtimer"
	"testing"
	"time"
)

func TestExampleAdd(t *testing.T) {
	now := time.Now()
	interval := 3 * time.Second
	gtimer.AddTimes(interval, 3, func() {
		fmt.Println(time.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
		now = time.Now()
	})
	select {}
}
