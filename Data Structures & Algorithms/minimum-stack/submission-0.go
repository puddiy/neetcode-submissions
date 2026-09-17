type MinStack struct {
	minStack []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		minStack: []int{},
		stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	
	minVal := val
	if len(this.minStack) > 0 {
		if top := this.minStack[len(this.minStack)-1]; top < minVal {
			minVal = top
		}
	}
	this.minStack = append(this.minStack, minVal)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
