package timeTick

import (
	"fmt"
	"github.com/gogf/gf/g/os/gtimer"
	"testing"
	"time"
)

func TestExampleAdd(t *testing.T) {
	now := time.Now()
	interval := 1400 * time.Millisecond
	gtimer.Add(interval, func() {
		fmt.Println(time.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
		now = time.Now()
	})
	select {}
}
