package main

import (
	"fmt"
	"os"
)

type Version struct {
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

type Releases struct {
	ReleaseName string
	Date        string
	Duration    int
	Platforms   []Platform
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("compute <path-to-data>")
		return
	}
}
