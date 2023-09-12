package main

import (
	"testing"
)

func TestHandleLine(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		hasSpeaker   bool
		expectedLine string
	}{
		{
			name:         "empty string",
			input:        "",
			hasSpeaker:   false,
			expectedLine: "",
		},
		{
			name:         "line number",
			input:        "202",
			hasSpeaker:   false,
			expectedLine: "",
		},
		{
			name:         "time",
			input:        "00:36:17.540 --> 00:36:38.369",
			hasSpeaker:   false,
			expectedLine: "",
		},
		{
			name:         "with speaker",
			input:        "@maxbeizer: some such nonsense",
			hasSpeaker:   false,
			expectedLine: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			speaker, result := HandleLine(tt.input)
			if (speaker != "" && !tt.hasSpeaker) || (speaker == "" && tt.hasSpeaker) {
				t.Errorf("expected %t, but got %q", tt.hasSpeaker, speaker)
			}
			if result != tt.expectedLine {
				t.Errorf("expected %q, but got %q", tt.expectedLine, result)
			}
		})
	}
}
