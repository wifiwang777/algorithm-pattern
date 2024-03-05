package sort

import (
	"math/rand"
	"slices"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	var items []int
	for i := 0; i < 10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomNumber := r.Intn(10) + 1
		items = append(items, randomNumber)
	}
	t.Log(items)
	QuickSort(items)
	t.Log(items)

	slices.Sort(items)
}
