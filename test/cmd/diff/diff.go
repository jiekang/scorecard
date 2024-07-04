package main

import (
	"encoding/json"
	"fmt"
	"os"

	dt "github.com/smlambert/scorecard/test/lib"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("diff <previous-release> <next-release>")
		return
	}

	prevFile := os.Args[1]
	nextFile := os.Args[2]

	data1, err := os.ReadFile(prevFile)
	if err != nil {
		fmt.Println("Unable to read: ", prevFile)
		panic(err)
	}

	data2, err := os.ReadFile(nextFile)
	if err != nil {
		fmt.Println("Unable to read: ", nextFile)
		panic(err)
	}

	var prevRelease dt.Release
	var nextRelease dt.Release
	json.Unmarshal(data1, &prevRelease)
	json.Unmarshal(data2, &nextRelease)

	compare(prevRelease, nextRelease)
}

func compare(prevRelease, nextRelease dt.Release) {
	panic("unimplemented")
}
