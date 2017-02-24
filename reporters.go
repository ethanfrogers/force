package main

import (
	"bytes"
	"strconv"
)

func ReporterJunit(result TestCoverage) string {

	var buf bytes.Buffer
	buf.WriteString(`<?xml version="1.0" encoding="UTF-8"?><testsuites>`)

	sortedCases := make(map[string][]string)

	for index := range result.SMethodNames {
		className := result.SClassNames[index]

		if _, ok := sortedCases[className]; !ok {
			sortedCases[className] = []string{}
			// continue
		}

		sortedCases[className] = append(sortedCases[className], result.SMethodNames[index])
	}

	for className, methods := range sortedCases {
		numMethods := strconv.Itoa(len(methods))
		buf.WriteString(`<testsuite name="` + className + `" tests="` + numMethods + `">`)
		for _, methodName := range methods {
			buf.WriteString(`<testcase classname="` + className + `" name="` + methodName + `"/>`)
		}
		buf.WriteString(`</testsuite>`)
	}

	buf.WriteString(`</testsuites>`)

	return buf.String()
}

func ReporterStandard(result TestCoverage) string {
	var buf bytes.Buffer
	var percent string
	buf.WriteString("Coverage:\n")
	for index := range result.NumberLocations {
		if result.NumberLocations[index] != 0 {
			locations := float64(result.NumberLocations[index])
			notCovered := float64(result.NumberLocationsNotCovered[index])
			percent = strconv.Itoa(int((locations-notCovered)/locations*100)) + "%"
		} else {
			percent = "0%"
		}
		buf.WriteString("  " + percent + "\t" + result.Name[index] + "\n")
	}
	buf.WriteString("\n\n")
	buf.WriteString("Results:\n")
	for index := range result.SMethodNames {
		buf.WriteString("  [PASS]  " + result.SClassNames[index] + "::" + result.SMethodNames[index] + "\n")
	}

	for index := range result.FMethodNames {
		buf.WriteString("  [FAIL]  " + result.FClassNames[index] + "::" + result.FMethodNames[index] + ": " + result.FMessage[index] + "\n")
		buf.WriteString("    " + result.FStackTrace[index] + "\n")
	}
	buf.WriteString("\n\n")
	return buf.String()
}
