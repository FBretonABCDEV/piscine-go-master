package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Tail == nil {
		l.Head = n
		l.Tail = n
		return
	}

	l.Tail.Next = n
	l.Tail = n
}