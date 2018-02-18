package main

import (
	"aueb.gr/cslabs/scheduler/model"
	"os"
	"aueb.gr/cslabs/scheduler/parser"
	"aueb.gr/cslabs/scheduler/algorithm"
	"container/heap"
	"aueb.gr/cslabs/scheduler/fitness"
	"time"
	"math/rand"
	"fmt"
	"strconv"
	"aueb.gr/cslabs/scheduler/custom_rules"
	"aueb.gr/cslabs/scheduler/output"
	"flag"
)

var generateFlag = flag.Bool("generate", false, "generate a schedule")

var title = flag.String("title", "", "represents the title for the schedule")
var preferencesFile = flag.String("prefs", "", "points to the file that contains the preferences in a CSV format")

func main() {
	flag.Parse()

	if *generateFlag {
		generate()
	} else {
		panic("No operation requested! Exiting.")
	}
}

func generate() {
	//Generate randomizer
	seed := rand.NewSource(time.Now().UnixNano())
	algorithm.Generator = rand.New(seed)

	if preferencesFile == nil || *preferencesFile == "" {
		panic("You did not provide a preferences file! Exiting.")
	}
	if title == nil || *title == "" {
		panic("You did not provide a title! Exiting.")
	}

	//Parse preferences file
	f, err := os.Open(*preferencesFile)
	if err != nil {
		panic(err.Error())
	}
	admins := parser.ParsePreferenceCSV(f, 5, 6)

	//Create times that are to be filled
	var times []model.DayTime
	for day := model.FirstDay; day <= model.LastDay; day++ {
		for hour := model.FirstHour; hour <= model.LastHour; hour++ {
			times = append(times, model.DayTime{Day: day, Time: hour})
		}
	}

	//Log time start
	timeStart := time.Now()

	//Initializing schedule generator and
	sampleSize := 100000
	model.CustomBlockRule = custom_rules.CustomBlockRules
	fmt.Println("Generating random schedules...")

	pq := make(algorithm.PriorityQueue, sampleSize)
	for i := 0; i < sampleSize; i++ {
		schedule := algorithm.GenerateRandomSchedule(admins, times)
		schedule.Index = i
		schedule.Fitness = scorer.CalculateFitness(schedule, admins, times)
		pq[i] = &schedule

		if i % 10000 == 0 && i != 0 {
			fmt.Println("Generated " + strconv.Itoa(i) + " random schedules...")
		}
	}
	fmt.Println("Generated " + strconv.Itoa(sampleSize) + " random schedules!\n")
	heap.Init(&pq)

	//Generate children until heap size < 5
	gen := 0
	for ;sampleSize > 5; {
		pq, sampleSize = algorithm.GenerateNextHeap(times, admins, pq, sampleSize)
		fmt.Print("Generated: \t" + strconv.Itoa(gen) + " gen \t")
		bestNow := pq[pq.Len() - 1]
		fmt.Println("(Best score now: " + strconv.Itoa(bestNow.Fitness)+ ")")
		gen += 1
	}

	//Retrieve the best
	best := heap.Pop(&pq).(*model.Schedule)
	bestSchedule := *best

	//Save as PDF (or HTML if that fails) and JSON
	err = output.GeneratePDF(*title, bestSchedule, admins, times, 5)
	if err != nil {
		fmt.Println(err.Error())
		output.GenerateHtml(*title, bestSchedule, admins, times, 5)
	}
	output.GenerateJson(*title, bestSchedule)

	fmt.Println("\nSchedule generated in " + strconv.Itoa(int(time.Since(timeStart).Seconds())) + " seconds!")
}