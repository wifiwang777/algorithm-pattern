package datastruct

type TreeNode[T comparable] struct {
	item  T
	left  *TreeNode[T]
	right *TreeNode[T]
}
