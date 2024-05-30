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
	fmt.Printf("Adding data for build group: %s\n", id)
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

		addPlatformData(platform, item)

		version := item["version"].(string)
		addVersionData(platform, version, item)
	}
}

func addPlatformData(platform string, data map[string]any) {
	fmt.Printf("Adding platform data for: %s\n", platform)
	if _, f := result.Platforms[platform]; !f {
		p := Platform{}
		p.Platform = platform
		p.Versions = make(map[string]Version)
		p.TestTargetTotals = TestTotals{0, 0, 0, 0, 0, 0}
		p.Duration = 0
		result.Platforms[platform] = p
	}
	p := result.Platforms[platform]
	p.Duration += int(data["buildDuration"].(float64))
	addTestTotals(&p.TestTargetTotals, data)
	result.Platforms[platform] = p
}

func addVersionData(platform string, version string, data map[string]any) {
	fmt.Printf("Adding version data for: %s %s\n", platform, version)
	p := result.Platforms[platform]
	if _, f := p.Versions[version]; f {
		v := Version{}
		v.Version = version
		v.TestTargetTotals = TestTotals{0, 0, 0, 0, 0, 0}
		v.Duration = 0
		p.Versions[version] = v
	}
	v := p.Versions[version]
	v.Duration += int(data["buildDuration"].(float64))
	addTestTotals(&v.TestTargetTotals, data)
	p.Versions[version] = v
}

func addTestTotals(totals *TestTotals, data map[string]any) {
	if data["testSummary"] != nil {
		summary := data["testSummary"].(map[string]any)
		totals.Disabled += int(summary["disabled"].(float64))
		totals.Executed += int(summary["executed"].(float64))
		totals.Failed += int(summary["failed"].(float64))
		totals.Passed += int(summary["passed"].(float64))
		totals.Skipped += int(summary["skipped"].(float64))
		totals.Total += int(summary["total"].(float64))
	}
}
