package main

import (
	"encoding/json"
	"fmt"
	"os"

	dt "github.com/smlambert/scorecard/test/lib"
)

var dataPath = "./data"
var releaseResult dt.Release

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

	releaseResult = dt.Release{ReleaseName: name, Date: date, Duration: 0, Platforms: make(map[string]*dt.Platform)}

	file := dataPath + "/builds.json"
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Unable to read: ", file)
		panic(err)
	}

	var builds dt.Builds
	json.Unmarshal(data, &builds)

	for _, id := range builds.Ids {
		computeBuild(id)
	}

	jsonOutput, _ := json.Marshal(releaseResult)
	filename := "data/" + name + ".json"
	os.WriteFile(filename, jsonOutput, 0644)
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

	computeData(buildData)
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
	computeData(testData)
}

func computeData(data []interface{}) {
	for _, i := range data {
		data := i.(map[string]any)

		releaseResult.Duration += int(data["buildDuration"].(float64))

		platform := data["platform"].(string)
		addPlatformData(platform, data)

		version := data["version"].(string)
		addVersionData(platform, version, data)
	}
}

func addPlatformData(platform string, data map[string]any) {
	fmt.Printf("Adding platform data for: %s\n", platform)
	if _, f := releaseResult.Platforms[platform]; !f {
		p := dt.Platform{}
		p.Platform = platform
		p.Versions = make(map[string]*dt.Version)
		p.TestTargetTotals = dt.TestTotals{Total: 0, Executed: 0, Passed: 0, Failed: 0, Disabled: 0, Skipped: 0}
		p.Duration = 0
		releaseResult.Platforms[platform] = &p
	}
	p := releaseResult.Platforms[platform]
	p.Duration += int(data["buildDuration"].(float64))
	addTestTotals(&p.TestTargetTotals, data)
}

func addVersionData(platform string, version string, data map[string]any) {
	fmt.Printf("Adding version data for: %s %s\n", platform, version)
	p := releaseResult.Platforms[platform]
	if _, f := p.Versions[version]; !f {
		v := dt.Version{}
		v.Version = version
		v.TestTargetTotals = dt.TestTotals{Total: 0, Executed: 0, Passed: 0, Failed: 0, Disabled: 0, Skipped: 0}
		v.Duration = 0
		p.Versions[version] = &v
	}
	v := p.Versions[version]
	v.Duration += int(data["buildDuration"].(float64))
	addTestTotals(&v.TestTargetTotals, data)
}

func addTestTotals(totals *dt.TestTotals, data map[string]any) {
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
