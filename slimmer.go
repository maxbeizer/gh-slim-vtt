package main

import (
	"regexp"
	"strconv"
	"strings"
)

var voiceTagRegex = regexp.MustCompile(`<v\s+([^>]+)>(.*)`)
var cueIdRegex = regexp.MustCompile(`^[a-f0-9-]+/\d+-\d+$`)

func HandleLine(currentSpeaker, line string) (speaker, processedLine string) {
	switch {
	case isHeader(line):
		return "", ""
	case isNumbersOnly(line):
		return "", ""
	case isCueId(line):
		return "", ""
	case isTimestmap(line):
		return "", ""
	default:
		return parseSpeakerLine(currentSpeaker, line)
	}
}

func parseSpeakerLine(currentSpeaker, line string) (string, string) {
	// Teams-style voice tags: <v Speaker>text</v>
	if matches := voiceTagRegex.FindStringSubmatch(line); matches != nil {
		speaker := strings.TrimSpace(matches[1])
		text := strings.TrimSpace(stripVoiceClosingTag(matches[2]))
		if text == "" {
			return "", ""
		}
		if speaker == currentSpeaker {
			return "", text
		}
		return speaker, " " + text
	}

	// Teams continuation line (no <v> tag but has </v>, or no colon at all)
	if strings.Contains(line, "</v>") || !strings.Contains(line, ":") {
		text := strings.TrimSpace(stripVoiceClosingTag(line))
		if text != "" {
			return "", text
		}
		return "", ""
	}

	// Zoom-style "Speaker: text"
	sp := strings.Split(line, ":")
	extractedSpeaker, extractedLine := sp[0], sp[1:]

	if extractedSpeaker == currentSpeaker {
		return "", strings.TrimLeft(strings.Join(extractedLine, " "), " ")
	}

	return extractedSpeaker, strings.Join(extractedLine, " ")
}

func stripVoiceClosingTag(s string) string {
	return strings.TrimSuffix(s, "</v>")
}

func isHeader(line string) bool {
	return strings.Contains(line, "WEBVTT")
}

func isNumbersOnly(line string) bool {
	if _, err := strconv.Atoi(line); err == nil {
		return true
	}
	return false
}

func isCueId(line string) bool {
	return cueIdRegex.MatchString(line)
}

func isTimestmap(line string) bool {
	pattern := `^\d{2}:\d{2}:\d{2}\.\d{3} --> \d{2}:\d{2}:\d{2}\.\d{3}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(line)
}
