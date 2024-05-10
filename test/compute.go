package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TestTarget struct {
	Name     string
	Duration int
}

type Version struct {
	Version          string
	TestTargetTotals TestTotals
	TestTotals       TestTotals
	Duration         int
	TestTargets      []TestTarget
}

type MachinesAvailable struct {
	DryRun  int
	Release int
}

type TestTotals struct {
	Total    int
	Executed int
	Passed   int
	Failed   int
	Disabled int
	Skipped  int
}

type Platform struct {
	Platform          string
	TestTargetTotals  TestTotals
	TestTotals        TestTotals
	MachinesAvailable MachinesAvailable
	Duration          int
	Versions          []Version
}

type Release struct {
	ReleaseName string
	Date        string
	Duration    int
	Platforms   []Platform
}

type Builds struct {
	Ids []string
}

var versions = []int{8, 11, 17, 21}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("compute <release-name> <date> [path-to-data] ")
		return
	}

	name := os.Args[1]
	date := os.Args[2]

	dataPath := "./data"
	if len(os.Args) > 3 {
		dataPath = os.Args[3]
	}

	result := Release{name, date, 0, []Platform{}}
	_ = result

	file := dataPath + "/builds.json"
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Unable to read: ", file)
		panic(err)
	}

	var builds Builds
	json.Unmarshal(data, &builds)

	for _, id := range builds.Ids {
		computeBuild(id)
	}

}

func computeBuild(id string) {
	fmt.Println("id ", id)
}
