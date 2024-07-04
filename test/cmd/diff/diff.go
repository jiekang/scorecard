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

	prevData, err := os.ReadFile(prevFile)
	if err != nil {
		fmt.Println("Unable to read: ", prevFile)
		panic(err)
	}

	nextData, err := os.ReadFile(nextFile)
	if err != nil {
		fmt.Println("Unable to read: ", nextFile)
		panic(err)
	}

	var prevRelease dt.Release
	var nextRelease dt.Release
	err = json.Unmarshal(prevData, &prevRelease)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(nextData, &nextRelease)
	if err != nil {
		panic(err)
	}

	result := nextRelease.Diff(prevRelease)
	jsonOutput, _ := json.Marshal(result)
	filename := "data/" + prevRelease.ReleaseName + "-" + nextRelease.ReleaseName + ".json"
	fmt.Println("Writing diff to file: " + filename)
	os.WriteFile(filename, jsonOutput, 0644)
}
