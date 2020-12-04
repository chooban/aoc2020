package io

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadFileLines(filepath string) (lines []string) {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	lines = strings.Split(string(content), "\n")
	return
}
