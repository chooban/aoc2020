package io

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadFileLines(filepath string) (contents []string) {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")

	contents = []string{}
	for _, line := range lines {
		if len(line) != 0 {
			// contents[i] = line
			contents = append(contents, line)
		}
	}
	return
}
