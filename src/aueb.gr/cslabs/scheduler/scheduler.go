package main

import (
	"aueb.gr/cslabs/scheduler/model"
	"os"
	"aueb.gr/cslabs/scheduler/parser"
	"aueb.gr/cslabs/scheduler/generator"
	"container/heap"
	"aueb.gr/cslabs/scheduler/fitness"
	"io/ioutil"
	"time"
	"math/rand"
	"fmt"
	"strconv"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
	generator.Generator = rand.New(s1)

	var times []model.DayTime
	for day := model.FirstDay; day <= model.LastDay; day++ {
		for hour := model.FirstHour; hour <= model.LastHour; hour++ {
			times = append(times, model.DayTime{Day: day, Time: hour})
		}
	}

	f, err := os.Open("test_schedule.csv")
	if err != nil {
		panic(err.Error())
	}
	admins := parser.ReadFromFile(f, 5, 6)
	sampleSize := 200000
	model.CustomBlockRule = customBlockRules

	pq := make(model.PriorityQueue, sampleSize)
	for i := 0; i < sampleSize; i++ {
		schedule := generator.GenerateRandomSchedule(admins, times)
		schedule.Index = i
		schedule.Fitness = scorer.CalculateFitness(schedule, admins, times)
		pq[i] = &schedule
		if i % 10000 == 0 {
			fmt.Println("Generated " + strconv.Itoa(i) + " random schedules...")
		}
	}
	fmt.Println("Generated " + strconv.Itoa(sampleSize) + " random schedules!")
	heap.Init(&pq)

	gen := 0
	for ;sampleSize > 5; {
		pq, sampleSize = generateHeapFromChildren(times, admins, pq, sampleSize)
		fmt.Print("Generated the n" + strconv.Itoa(gen) + " generation! ")
		bestNow := pq[pq.Len() - 1]
		fmt.Println("(Best score now: " + strconv.Itoa(bestNow.Fitness)+ ")")
		gen += 1
	}

	best := heap.Pop(&pq).(*model.Schedule)
	bestSchedule := *best

	html := generator.GenerateHtml(bestSchedule, admins, times)
	err = ioutil.WriteFile("schedule.html", []byte(html), 0644)
	if err != nil {
		panic(err)
	}
}

func generateHeapFromChildren(times []model.DayTime, admins []model.Admin, pqOld model.PriorityQueue, currentSize int) (model.PriorityQueue, int) {
	matingSize := int(float64(currentSize) * 0.6)
	staticMatingSize := matingSize
	if matingSize % 2 == 1 {
		matingSize--
	}

	newSize := matingSize + matingSize / 2
	pq := make(model.PriorityQueue, newSize)
	index := 0

	for i := 0; i < staticMatingSize; i += 2 {
		if matingSize < 2 {
			break
		}
		s1Index := generator.Generator.Intn(matingSize)
		s2Index := generator.Generator.Intn(matingSize)
		s1 := heap.Remove(&pqOld, s1Index).(*model.Schedule)
		s2 := heap.Remove(&pqOld, s2Index).(*model.Schedule)
		sChild := generator.MateSchedules(times, *s1, *s2)
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