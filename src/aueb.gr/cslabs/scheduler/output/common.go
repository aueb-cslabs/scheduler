package output

import (
	"os"
	_ "fmt"
	"time"
	"strconv"
)

func prepareOutDir() {
	if _, err := os.Stat("out"); os.IsNotExist(err) {
		os.Mkdir("out", 0644)
	}
}

func getOutputFile(ext string) string {
	t := time.Now()
	prefix := "out/out_" + t.Format("20060102") + "_"
	for i := 1;; i++ {
		name := prefix + strconv.Itoa(i) + "." + ext
		if _, err := os.Stat(name); os.IsNotExist(err) {
			return name
		}
	}
	return ""
}