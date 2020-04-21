package collection

// 这是一个栈的简单结构
type Stack struct {
	data []string
}

func (s *Stack) Push(item string) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() string {
	length := len(s.data)
	popItem := s.data[length-1]
	s.data[length-1] = ""
	s.data = s.data[:length-1]
	return popItem
}

func (s *Stack) Size() int {
	return len(s.data)
}

func f(i, j, k int) int {
	l:=10
	println(l)
	return 0
}

func test() {

	type binOp func(int,int) int

	var op binOp
	add := func(i,j int) int {return i+ j}
	op = add
	n := op(10, 20)
	println(n)
}
