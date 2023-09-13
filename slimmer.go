package main

import (
	"regexp"
	"strconv"
	"strings"
)

func HandleLine(currentSpeaker, line string) (speaker, processedLine string) {
	switch {
	case isHeader(line):
		return "", ""
	case isNumbersOnly(line):
		return "", ""
	case isTimestmap(line):
		return "", ""
	default:
		sp := strings.Split(line, ":")
		extractedSpeaker, extractedLine := sp[0], sp[1:]

		if extractedSpeaker == currentSpeaker {
			return "", strings.TrimLeft(strings.Join(extractedLine, " "), " ")
		}

		return extractedSpeaker, strings.Join(extractedLine, " ")
	}
}

func isHeader(line string) bool {
	// Check if the line is a header
	return strings.Contains(line, "WEBVTT")
}

func isNumbersOnly(line string) bool {
	// Check if the line is a number
	if _, err := strconv.Atoi(line); err == nil {
		return true
	}
	return false
}

func isTimestmap(line string) bool {
	// Define the regular expression pattern
	pattern := `^\d{2}:\d{2}:\d{2}\.\d{3} --> \d{2}:\d{2}:\d{2}\.\d{3}$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Check if the line matches the regular expression
	return regex.MatchString(line)
}
