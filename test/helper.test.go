package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var failedTests int

func reportResult[T any](name string, records []T, err error) {
	if err != nil {
		failedTests++
		fmt.Printf("%s: failed: %s\n", name, err.Error())
		return
	}

	fmt.Printf("%s: ok (%d result(s))\n", name, len(records))
	if len(records) == 0 {
		return
	}

	encoded, err := json.MarshalIndent(records[0], "", "  ")
	if err != nil {
		failedTests++
		fmt.Printf("%s: failed to encode first result: %s\n", name, err.Error())
		return
	}

	fmt.Printf("%s: first result\n%s\n", name, encoded)
}

func currentWeek() string {
	year, week := time.Now().ISOWeek()
	return fmt.Sprintf("%04d%02d", year, week)
}
