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
	"aueb.gr/cslabs/scheduler/output"
	"flag"
	"encoding/json"
	"io/ioutil"
)

//TODO Here you can import your custom rule set
import "aueb.gr/cslabs/scheduler/custom_rules"

var generateFlag = flag.Bool("generate", false, "generate a schedule")
var docsFlag = flag.String("docs", "", "generate docs from existing")

var title = flag.String("title", "", "represents the title for the schedule")
var preferencesFile = flag.String("prefs", "", "points to the file that contains the preferences in a CSV format")
var configFile = flag.String("config", "config.json", "points to the config.json file")

func main() {

	flag.Parse()
	config, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(config, &model.Config)
	if err != nil {
		panic(err.Error())
	}

	if *generateFlag {
		generate()
	} else if *docsFlag != "" {
		docs()
	} else {
		panic("No operation requested! Exiting.")
	}
}

func loadPrefsTimes() ([]model.DayHour, []model.Admin, int) {
	f, err := os.Open(*preferencesFile)
	if err != nil {
		panic(err.Error())
	}
	admins := parser.ParsePreferenceCSV(f, model.Config.PreferencesDays, model.Config.PreferencesDayLength)

	//Create times that are to be filled
	totalHours := 0
	var times []model.DayHour
	for day := model.Config.ScheduleFirstDay; day <= model.Config.ScheduleLastDay; day++ {
		dayIgnored := intInSlice(day, model.Config.IgnoreDays)
		for hour := model.Config.ScheduleFirstHour; hour <= model.Config.ScheduleLastHour; hour++ {
			dayTime := model.DayHour{Day: day, Time: hour}
			ignored := dayIgnored || intInSlice(hour, model.Config.IgnoreHours) || dayTimeInSlice(dayTime, model.Config.IgnoreDayTimes)
			dayTime.Ignored = ignored
			times = append(times, dayTime)
			if !ignored {
				totalHours++
			}
		}
	}
	return times, admins, totalHours
}

func docs() {
	dat, err := ioutil.ReadFile(*docsFlag)
	if err != nil {
		panic(err.Error())
	}
	schedule := model.Schedule{}
	json.Unmarshal(dat, &schedule)
	times, admins, _ := loadPrefsTimes()
	generateDocs(schedule, admins, times)

	fmt.Println("\nDocuments regenerated!")
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

	times, admins, totalHours := loadPrefsTimes()

	//Log time start
	timeStart := time.Now()

	//TODO Here you can specify the custom rules method
	model.CustomBlockRule = custom_rules.CustomBlockRules

	//Initializing schedule generator
	sampleSize := 150000
	fitness.HoursPerAdmin = len(admins) / totalHours
	fmt.Println("Generating random schedules...")

	pq := make(algorithm.PriorityQueue, sampleSize)
	for i := 0; i < sampleSize; i++ {
		schedule := algorithm.GenerateRandomSchedule(admins, times)
		schedule.Index = i
		schedule.Fitness = fitness.CalculateFitness(schedule, admins, times)
		pq[i] = &schedule

		if i % 10000 == 0 && i != 0 {
			fmt.Println("Generated " + strconv.Itoa(i) + " random schedules...")
		}
	}
	fmt.Println("Generated " + strconv.Itoa(sampleSize) + " random schedules!\n")
	heap.Init(&pq)

	//Generate children until heap size < 5
	gen := 1
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
	bestSchedule.Title = *title

	//Save as PDF (or HTML if that fails) and JSON
	generateDocs(bestSchedule, admins, times)

	fmt.Println("\nSchedule generated in " + strconv.Itoa(int(time.Since(timeStart).Seconds())) + " seconds!")
}

func generateDocs(schedule model.Schedule, admins []model.Admin, times []model.DayHour) {
	err := output.GeneratePDF(schedule, admins, times, model.Config.ScheduleDayLength())
	if err != nil {
		fmt.Println(err.Error())
		output.GenerateHtml(schedule, admins, times)
	}
	err = output.GenerateOfficialPDF(schedule, admins, times, model.Config.ScheduleDayLength())
	if err != nil {
		fmt.Println(err.Error())
		output.GenerateOfficialHtml(schedule, admins, times)
	}
	output.GenerateJson(schedule)
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func dayTimeInSlice(a model.DayHour, list []model.DayHour) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}