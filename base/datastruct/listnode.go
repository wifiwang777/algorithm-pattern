package datastruct

type ListNode[T any] struct {
	item T
	next *ListNode[T]
}
