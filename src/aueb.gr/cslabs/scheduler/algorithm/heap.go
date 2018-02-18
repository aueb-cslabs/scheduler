package algorithm

import (
	"container/heap"
	"aueb.gr/cslabs/scheduler/fitness"
	"aueb.gr/cslabs/scheduler/model"
)

func GenerateNextHeap(times []model.DayTime, admins []model.Admin, pqOld PriorityQueue, currentSize int) (PriorityQueue, int) {
	matingSize := int(float64(currentSize) * 0.6)
	staticMatingSize := matingSize
	if matingSize % 2 == 1 {
		matingSize--
	}

	newSize := matingSize + matingSize / 2
	pq := make(PriorityQueue, newSize)
	index := 0

	for i := 0; i < staticMatingSize; i += 2 {
		if matingSize < 2 {
			break
		}
		s1Index := Generator.Intn(matingSize)
		s2Index := Generator.Intn(matingSize)
		s1 := heap.Remove(&pqOld, s1Index).(*model.Schedule)
		s2 := heap.Remove(&pqOld, s2Index).(*model.Schedule)
		sChild := MateSchedules(times, *s1, *s2)
		sChild.Fitness = scorer.CalculateFitness(sChild, admins, times)

		s1.Index = index
		pq[index] = s1
		index++
		s2.Index = index
		pq[index] = s2
		index++
		sChild.Index = index
		pq[index] = &sChild
		index++
		matingSize -= 2
	}
	heap.Init(&pq)

	return pq, newSize
}
