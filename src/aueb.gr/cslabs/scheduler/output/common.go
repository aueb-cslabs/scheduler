package output

import (
	_ "fmt"
	"os"
	"strconv"
	"time"
)

func prepareOutDir() {
	if _, err := os.Stat("out"); os.IsNotExist(err) {
		os.Mkdir("out", 0644)
	}
}

func getOutputFile(pref string, ext string) string {
	t := time.Now()
	prefix := "out/" + pref + "_" + t.Format("20060102") + "_"
	for i := 1; ; i++ {
		name := prefix + strconv.Itoa(i) + "." + ext
		if _, err := os.Stat(name); os.IsNotExist(err) {
			if i == 1 {
				return "out/" + pref + "_" + t.Format("20060102") + "." + ext
			}
			return name
		}
	}
	return ""
}
