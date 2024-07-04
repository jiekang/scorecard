package data

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
	Versions          map[string]*Version
}

type Release struct {
	ReleaseName string
	Date        string
	Duration    int
	Platforms   map[string]*Platform
}

type Builds struct {
	Ids []string
}
