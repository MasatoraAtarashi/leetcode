package main

type MinStack struct {
	main Stack[int]
	min  Stack[int]
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.main.Push(val)
	if this.min.Empty() || val <= this.min.Top() {
		this.min.Push(val)
	}
}

func (this *MinStack) Pop() {
	if this.main.Top() == this.min.Top() {
		this.min.Pop()
	}
	this.main.Pop()
}

func (this *MinStack) Top() int {
	return this.main.Top()
}

func (this *MinStack) GetMin() int {
	return this.min.Top()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
