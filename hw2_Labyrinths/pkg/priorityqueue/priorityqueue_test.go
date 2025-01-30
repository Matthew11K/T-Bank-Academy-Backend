package priorityqueue_test

import (
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_2-go-Matthew11K/pkg/priorityqueue"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPriorityQueue(t *testing.T) {
	pq := priorityqueue.NewPriorityQueue()
	item1 := &priorityqueue.Item{
		Coordinate: domain.Coordinate{Row: 0, Col: 0},
		Priority:   1.0,
	}
	item2 := &priorityqueue.Item{
		Coordinate: domain.Coordinate{Row: 1, Col: 1},
		Priority:   0.5,
	}

	pq.Push(item1)
	pq.Push(item2)

	require.Equal(t, 2, pq.Len())

	popped := pq.Pop().(*priorityqueue.Item)
	assert.Equal(t, domain.Coordinate{Row: 1, Col: 1}, popped.Coordinate)

	require.Equal(t, 1, pq.Len())

	popped = pq.Pop().(*priorityqueue.Item)
	assert.Equal(t, domain.Coordinate{Row: 0, Col: 0}, popped.Coordinate)

	require.Equal(t, 0, pq.Len())
}
