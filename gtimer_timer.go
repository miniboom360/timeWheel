package timeTick

import (
	"tests/timeTick/gtype"
	"time"
)

type Timer struct {
	status *gtype.Int			//定时器状态
	wheels []*wheel		//分层时间轮对象
	length int			//分层层数
	number int			//每一层Slot Number
	intervalMs int64	//最小时间刻度(毫秒)
}

//单层时间轮
type wheel struct {
	timer *Timer		//所属定时器
	level int			//所属分层索引号
	slots []*gtype.List	//所有的循环任务项，按照Slot Number进行分组
	number int64		//Slot Number=len(slots)
	ticks *gtype.Int64	//当前时间轮已转动的刻度数量
	totalMs int64		//整个时间轮的时间长度(毫秒)=number*interval ???
	createMs int64
	intervalMs int64
}

func New(slot int, interval time.Duration, level...int)*Timer{
	length := DEFAULT_WHEEL_LEVEL
	if len(level) > 0 {
		length = level[0]
	}
	t := &Timer{
		status:gtype.NewInt(STATUS_RUNNING),
		wheels : make([]*wheel, length),
		length:length,
		number:slot,
		intervalMs:interval.Nanoseconds()/1e6,
	}

	for i := 0; i <length; i++{
		if i > 0 {
			n := time.Duration(t.wheels[i - 1].totalMs) * time.Millisecond
			w := t

		}
	}

}

