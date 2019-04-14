package timeTick

import (
  "time"
  "timeWheel/gtype"
)

// 执行时间轮刻度逻辑
func (w *wheel) proceed(){
  n := w.ticks.Add(1)
  l := w.slots[int(n%w.number)]
  length := l.Len()
  if length > 0{
    go func(l *gtype.List, nowTicks int64) {
      entry := (*Entry)(nil)
      nowMs := time.Now().UnixNano()/1e6
      for i := length; i > 0; i--{
        if v := l.PopFront(); v == nil{
          break
        }else {
          entry = v.(*Entry)
        }
        // 是否满足运行条件
        runnable, addable := entry.check(nowTicks, nowMs)
        if runnable {
          //异步执行运行
          go func(entry *Entry) {
            defer func() {
              if err := recover(); err != nil{
                if err != gPANIC_EXIT{
                  panic(err)
                }else {
                  entry.Close()
                }
              }
              if entry.Status() == STATUS_RUNNING{  //todo:entry将running改为ready?
                entry.SetStatus(STATUS_READY)
              }
            }()
            entry.job()
          }(entry)
        }
        // 是否继续添运行，滚动任务
        if addable{
          //优先从chird time wheel开始添加
          entry.wheel.timer.doAddEntryByParent(entry.rawIntervalMs, entry)
        }
      }
    }(l,n)
  }
}
