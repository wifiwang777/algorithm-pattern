package datastruct

import "testing"

func TestStack(t *testing.T) {
	stack1 := new(Stack[int])
	stack1.Push(1)
	stack1.Push(2)
	stack1.Push(3)
	stack1.Push(4)
	for !stack1.Empty() {
		v := stack1.Pop()
		t.Log(v)
	}

	t.Log("\n ----------------- \n")

	stack2 := new(Stack[string])
	stack2.Push("1")
	stack2.Push("2")
	stack2.Push("3")
	stack2.Push("4")
	for !stack2.Empty() {
		v := stack2.Pop()
		t.Log(v)
	}

}
