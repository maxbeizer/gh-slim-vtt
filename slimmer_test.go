package main

import (
	"testing"
)

func TestHandleLine(t *testing.T) {
	tests := []struct {
		name           string
		line           string
		currentSpeaker string
		hasSpeaker     bool
		expectedLine   string
	}{
		{
			name:         "header",
			line:         "WEBVTT",
			hasSpeaker:   false,
			expectedLine: "",
		},
		{
			name:         "empty string",
			line:         "",
			hasSpeaker:   false,
			expectedLine: "",
		},
		{
			name:         "line number",
			line:         "202",
			hasSpeaker:   false,
			expectedLine: "",
		},
		{
			name:         "time",
			line:         "00:36:17.540 --> 00:36:38.369",
			hasSpeaker:   false,
			expectedLine: "",
		},
		{
			name:         "with speaker",
			line:         "@maxbeizer: some such nonsense",
			hasSpeaker:   true,
			expectedLine: " some such nonsense",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			speaker, result := HandleLine("", tt.line)
			if (speaker != "" && !tt.hasSpeaker) || (speaker == "" && tt.hasSpeaker) {
				t.Errorf("expected %t, but got %q", tt.hasSpeaker, speaker)
			}
			if result != tt.expectedLine {
				t.Errorf("expected %q, but got %q", tt.expectedLine, result)
			}
		})
	}
}
