package timeTick

import (
  "time"
  "timeWheel/gtype"
)

type Entry struct{
  wheel *wheel  //所属时间轮
  job JobFunc   //注册循环任务方法
  singleton *gtype.Bool //任务是否单例运行
  status *gtype.Int //任务状态(0: ready; 1:running; 2: stopped; -1:closed),层级entry共享状态
  times *gtype.Int //还需运行次数
  create int64  //注册时的时间轮tick
  interval int64 //设置的运行间隔(时间轮刻度数量)
  createMs int64 //创建时间(毫秒)
  intervalMs int64 //间隔时间(毫秒)
  rawIntervalMs int64 // 原始间隔
}

type JobFunc = func()

func (w *wheel) addEntry(interval time.Duration, job JobFunc, singleton bool, times int, status int) *Entry{
  ms := interval.Nanoseconds() / 1e6
  num := ms/w.intervalMs
  if num == 0{
    // 如果安装的任务间隔小于时间轮刻度，
    // 那么将会在下一刻度被执行
    num = 1
  }
  nowMs := time.Now().UnixNano()/1e6
  ticks = w.ticks.Val()
  entry := &Entry{
    wheel:w,
    job:job,
    times:gtype.NewInt(times),
    status:gtype.NewInt(status),
    create:ticks,
    interval:num,
    singleton:gtype.NewBool(singleton),
    createMs:nowMs,
    intervalMs:ms,
    rawIntervalMs:ms,
  }
  // 安装任务
  w.slots[(ticks + num) % w.number].PushBack(entry)
  return entry
}