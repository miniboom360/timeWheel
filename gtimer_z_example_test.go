package timeTick

import (
	"fmt"
	"github.com/gogf/gf/g/os/gtimer"
	"time"
)

func ExampleAdd() {
	now := time.Now()
	interval := 1400 * time.Millisecond
	gtimer.Add(interval, func() {
		fmt.Println(time.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
		now = time.Now()
	})
	select {}
}
