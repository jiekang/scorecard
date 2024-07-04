package data

type Duration int

func (s Duration) Diff(toCompare Duration) Duration {
	return s - toCompare
}

type TestTarget struct {
	Name     string
	Duration Duration
}

type Version struct {
	Version          string
	TestTargetTotals TestTotals
	TestTotals       TestTotals
	Duration         Duration
	TestTargets      map[string]TestTarget
}

type Versions map[string]*Version

func (s Versions) Diff(toCompare Versions) Versions {
	result := make(Versions)

	for key, value := range s {
		v := Version{}
		v.Version = value.Version
		v.TestTargetTotals = value.TestTargetTotals
		v.TestTotals = value.TestTotals
		v.Duration = value.Duration
		v.TestTargets = value.TestTargets

		result[key] = &v
		if val, f := toCompare[key]; f {
			v.TestTargetTotals = value.TestTargetTotals.Diff(val.TestTargetTotals)
			v.TestTotals = value.TestTotals.Diff(val.TestTotals)
			v.Duration = value.Duration.Diff(val.Duration)
			// TODO Compare TestTargets
			v.TestTargets = value.TestTargets
		}
	}

	return result
}

type MachinesAvailable struct {
	DryRun  int
	Release int
}

func (s MachinesAvailable) Diff(toCompare MachinesAvailable) MachinesAvailable {
	return MachinesAvailable{DryRun: s.DryRun - toCompare.DryRun, Release: s.Release - toCompare.Release}
}

type TestTotals struct {
	Total    int
	Executed int
	Passed   int
	Failed   int
	Disabled int
	Skipped  int
}

func (s TestTotals) Diff(toCompare TestTotals) TestTotals {
	result := TestTotals{}
	result.Total = s.Total - toCompare.Total
	result.Executed = s.Executed - toCompare.Executed
	result.Passed = s.Passed - toCompare.Passed
	result.Failed = s.Failed - toCompare.Failed
	result.Disabled = s.Disabled - toCompare.Disabled
	result.Skipped = s.Skipped - toCompare.Skipped

	return result
}

type Platform struct {
	Platform          string
	TestTargetTotals  TestTotals
	TestTotals        TestTotals
	MachinesAvailable MachinesAvailable
	Duration          Duration
	Versions          Versions
}

type Platforms map[string]*Platform

func (s Platforms) Diff(toCompare Platforms) Platforms {
	result := make(Platforms)
	for key, value := range s {
		p := Platform{}
		p.Platform = value.Platform
		p.TestTargetTotals = value.TestTargetTotals
		p.TestTotals = value.TestTotals
		p.MachinesAvailable = value.MachinesAvailable
		p.Duration = value.Duration
		p.Versions = value.Versions
		result[key] = &p

		if v, f := toCompare[key]; f {
			p.TestTargetTotals = value.TestTargetTotals.Diff(v.TestTargetTotals)
			p.TestTotals = value.TestTotals.Diff(v.TestTotals)
			p.MachinesAvailable = value.MachinesAvailable.Diff(v.MachinesAvailable)
			p.Duration = value.Duration.Diff(v.Duration)
			p.Versions = value.Versions.Diff(v.Versions)
		}
	}
	return result
}

type Release struct {
	ReleaseName      string
	Date             string
	Duration         Duration
	TestTargetTotals TestTotals
	Platforms        Platforms
}

type Builds struct {
	Ids []string
}

func (s Release) Diff(toCompare Release) Release {
	name := toCompare.ReleaseName + " " + s.ReleaseName
	return Release{ReleaseName: name, Date: "comparison", Duration: s.Duration.Diff(toCompare.Duration), Platforms: s.Platforms.Diff(toCompare.Platforms), TestTargetTotals: s.TestTargetTotals.Diff(toCompare.TestTargetTotals)}
}
