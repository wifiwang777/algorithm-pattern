package datastruct

import "testing"

func TestQueue(t *testing.T) {
	queue1 := new(Queue[int])
	queue1.Enqueue(1)
	queue1.Enqueue(2)
	queue1.Enqueue(3)
	queue1.Enqueue(4)
	for !queue1.Empty() {
		v := queue1.Dequeue()
		t.Log(v)
	}

	t.Log("\n ----------------- \n")

	queue2 := new(Queue[string])
	queue2.Enqueue("1")
	queue2.Enqueue("2")
	queue2.Enqueue("3")
	queue2.Enqueue("4")
	for !queue2.Empty() {
		v := queue2.Dequeue()
		t.Log(v)
	}

}
