package output

import (
	"aueb.gr/cslabs/scheduler/model"
	"encoding/json"
	"os"
	"bufio"
)

func GenerateJson(schedule model.Schedule) error {
	prepareOutDir()
	js, _ := json.MarshalIndent(schedule, "", "\t")
	f, err := os.Create(getOutputFile("json"))
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
