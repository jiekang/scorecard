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
	TestTargets      map[string]TestTarget
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
	Versions          map[string]Version
}

type Release struct {
	ReleaseName string
	Date        string
	Duration    int
	Platforms   map[string]Platform
}

type Builds struct {
	Ids []string
}

var dataPath = "./data"
var result Release

func main() {
	if len(os.Args) < 2 {
		fmt.Println("compute <release-name> <date> [path-to-data] ")
		return
	}

	name := os.Args[1]
	date := os.Args[2]

	if len(os.Args) > 3 {
		dataPath = os.Args[3]
	}

	result = Release{name, date, 0, make(map[string]Platform)}

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

	fmt.Printf("%v", result)

}

func computeBuild(id string) {
	fmt.Println("Adding data for build group: ", id)
	computeBuildData(id)
	computeTestData(id)
}

func computeBuildData(id string) {
	file := dataPath + "/child/jdk-" + id + "-compute.json"
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Unable to read: ", file)
		panic(err)
	}
	var buildData []interface{}
	json.Unmarshal(data, &buildData)
}

func computeTestData(id string) {
	file := dataPath + "/child/test-" + id + "-compute.json"
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Unable to read: ", file)
		panic(err)
	}
	var testData []interface{}
	json.Unmarshal(data, &testData)

	for _, i := range testData {
		item := i.(map[string]any)

		platform := item["platform"].(string)

		if _, f := result.Platforms[platform]; !f {
			p := Platform{}
			p.Platform = platform
			ttt := TestTotals{0, 0, 0, 0, 0, 0}
			if item["testSummary"] != nil {
				summary := item["testSummary"].(map[string]any)
				ttt.Disabled += int(summary["disabled"].(float64))
			}
			p.TestTargetTotals = ttt
			result.Platforms[platform] = p
		}
	}
}
