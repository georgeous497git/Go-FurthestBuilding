package furthest

import "container/heap"

type PriorityQueue []int

func (h PriorityQueue) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h PriorityQueue) Len() int {
	return len(h)
}

func (h *PriorityQueue) Pop() interface{} {
	original := *h
	l := len(*h)
	x := original[l-1]
	*h = original[:l-1]

	return x
}

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func FurthestBuilding(heights []int, bricks int, ladders int) int {

	pq := &PriorityQueue{}
	heap.Init(pq)

	maxLen := len(heights) - 1
	var next int

	for i, v := range heights {
		if maxLen >= i+1 {
			next = heights[i+1]
		} else {
			return i
		}

		if next > v {
			heap.Push(pq, next-v)

			if pq.Len() > ladders {
				min := heap.Pop(pq).(int)

				if min <= bricks {
					bricks = bricks - min
				} else {
					return i
				}
			}
		}
	}

	return -1
}
