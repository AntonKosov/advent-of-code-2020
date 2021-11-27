package aoc

type Stack struct {
	data []interface{}
}

func (l *Stack) Push(v interface{}) {
	old := l.data
	l.data = []interface{}{v}
	l.data = append(l.data, old...)
}

func (l *Stack) Pop() interface{} {
	v := l.data[0]
	l.data = l.data[1:]
	return v
}

func (l *Stack) Peek() interface{} {
	return l.data[0]
}

func (l *Stack) IsEmpty() bool {
	return len(l.data) == 0
}
