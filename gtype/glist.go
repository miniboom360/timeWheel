package gtype

import "container/list"

type (
	List struct {
		mu *RWMutex
		list *list.List
	}

	Element = list.Element
)

func New(unsafe...bool) *List{
	return &List{
		mu : NewRWMutex(unsafe...),
		list: list.New(),
	}
}

// 链表头入栈数据
func (l *List) PushFront(v interface{}) (e *Element) {
	l.mu.Lock()
	defer l.mu.Unlock()
	e = l.list.PushFront(v)
	return
}

// 批量从链表尾端出栈数据项(删除)
func (l *List) BatchPopBack(max int) (values []interface{}){
	l.mu.Lock()
	defer l.mu.Unlock()

	length := l.list.Len()
	if length > 0 {
		if max > 0 && max < length{
			length = max
		}
		tmp := (*Element)(nil)
		values = make([]interface{}, length)
		for i := 0; i < length; i++{
			tmp = l.list.Back()
			values[i] = l.list.Remove(tmp)
		}
	}
	return
}

func (l *List) PopBackAll() []interface{}{
	return l.BatchPopBack(-1)
}
