package priorityqueue

import "github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"

type Item struct {
	Coordinate domain.Coordinate
	Priority   float64
	Index      int
}

type PriorityQueue []*Item

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]

	return item
}
