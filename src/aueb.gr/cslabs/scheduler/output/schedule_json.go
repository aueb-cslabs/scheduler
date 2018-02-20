package output

import (
	"aueb.gr/cslabs/scheduler/model"
	"bufio"
	"encoding/json"
	"os"
)

func GenerateJson(schedule model.Schedule) error {
	prepareOutDir()
	js, _ := json.MarshalIndent(schedule, "", "\t")
	f, err := os.Create(getOutputFile("schedule", "json"))
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	_, err = w.Write(js)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}
