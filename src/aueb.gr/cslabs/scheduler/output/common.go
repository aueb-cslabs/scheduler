package output

import "os"

func prepareOutDir() {
	if _, err := os.Stat("out"); os.IsNotExist(err) {
		os.Mkdir("out", 0644)
	}
}
