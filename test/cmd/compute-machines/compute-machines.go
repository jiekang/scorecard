package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var dataFile = "./data/machine-results.json"

type Machine struct {
	Online  int
	Offline int
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("compute-machines")
		return
	}
	data, err := os.ReadFile(dataFile)
	if err != nil {
		fmt.Println("Unable to read: ", dataFile)
		panic(err)
	}
	var input []map[string]any
	err = json.Unmarshal(data, &input)
	if err != nil {
		panic(err)
	}
	osMap := make(map[string]map[string]*Machine)
	for _, machine := range input {
		os := machine["os"].(string)
		if _, f := osMap[os]; !f {
			osMap[os] = make(map[string]*Machine)
		}
		arch := machine["arch"].(string)
		archMap := osMap[os]

		if _, f := archMap[arch]; !f {
			archMap[arch] = &Machine{Online: 0, Offline: 0}
		}

		offline := machine["offline"].(bool)
		if offline {
			archMap[arch].Offline++
		} else {
			archMap[arch].Online++
		}
	}

	jsonOutput, _ := json.Marshal(osMap)
	filename := "data/machine-counts.json"
	os.WriteFile(filename, jsonOutput, 0644)
}
