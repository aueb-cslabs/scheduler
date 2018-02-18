package output

import (
	"aueb.gr/cslabs/scheduler/model"
	"encoding/json"
	"os"
	"bufio"
)

func GenerateJson(title string, schedule model.Schedule) error {
	prepareOutDir()
	js, _ := json.Marshal(schedule)
	f, err := os.Create("out/" + title + ".json")
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
