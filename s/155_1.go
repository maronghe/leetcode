package s

// time:16 ms	memery:7.6 MB	

// method 1
type record struct {
	val int
	min int
}

type MinStack struct {
	stack []record
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []record{},
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, record{x, x})
	} else {
		l := this.stack[len(this.stack)-1]
		if l.min <= x {
			this.stack = append(this.stack, record{x, l.min})
		} else {
			this.stack = append(this.stack, record{x, x})
		}
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}

// method 2

// time: 16 ms  memery:	7.5 MB

//type MinStack struct {
//	stack    []int
//	minStack []int // 辅助作用
//}

///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{
//		stack:    []int{},
//		minStack: []int{},
//	}
//}

//func (this *MinStack) Push(x int) {
//	if len(this.minStack) == 0 || x <= this.GetMin() {
//		this.minStack = append(this.minStack, x)
//	}
//	this.stack = append(this.stack, x)
//}

//func (this *MinStack) Pop() {
//	if this.minStack[len(this.minStack)-1] == this.stack[len(this.stack)-1] {
//		this.minStack = this.minStack[:len(this.minStack)-1]
//	}
//	this.stack = this.stack[:len(this.stack)-1]
//}

//func (this *MinStack) Top() int {
//	return this.stack[len(this.stack)-1]
//}

//func (this *MinStack) GetMin() int {
//	return this.minStack[len(this.minStack)-1]
//}


// method 3 
// time: 20 ms	memery: 7.8 MB

//import "math"

//type MinStack struct {
//	stack    []int
//	minStack []int // 辅助作用
//}

///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{
//		stack:    []int{},
//		minStack: []int{math.MaxInt64},
//	}
//}

//func (this *MinStack) Push(x int) {
//	this.stack = append(this.stack, x)
//	m := min(this.minStack[len(this.minStack)-1], x)
//	this.minStack = append(this.minStack, m)
//}

//func (this *MinStack) Pop() {
//	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}