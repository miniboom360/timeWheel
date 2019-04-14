package gtype

import "container/list"

type (
	List struct {
		mu   *RWMutex
		list *list.List
	}

	Element = list.Element
)

func New(unsafe ...bool) *List {
	return &List{
		mu:   NewRWMutex(unsafe...),
		list: list.New(),
	}
}

// 从链表头端出栈数据项(删除)
func (l *List) PopFront() (value interface{}) {
	l.mu.Lock()
	if e := l.list.Front(); e != nil {
		value = l.list.Remove(e)
	}
	l.mu.Unlock()
	return
}

// 获取链表长度
func (l *List) Len() (length int) {
	l.mu.RLock()
	length = l.list.Len()
	l.mu.RUnlock()
	return
}

// 链表头入栈数据
func (l *List) PushFront(v interface{}) (e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	e = l.list.PushFront(v)
	return
}

// 往链表尾入栈数据项
func (l *List) PushBack(v interface{}) (e *Element) {
	l.mu.Lock()
	e = l.list.PushBack(v)
	l.mu.Unlock()
	return
}

// 批量从链表尾端出栈数据项(删除)
func (l *List) BatchPopBack(max int) (values []interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	length := l.list.Len()
	if length > 0 {
		if max > 0 && max < length {
			length = max
		}
		tmp := (*Element)(nil)
		values = make([]interface{}, length)
		for i := 0; i < length; i++ {
			tmp = l.list.Back()
			values[i] = l.list.Remove(tmp)
		}
	}
	return
}

func (l *List) PopBackAll() []interface{} {
	return l.BatchPopBack(-1)
}
