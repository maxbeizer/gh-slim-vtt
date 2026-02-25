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
		// Teams VTT format
		{
			name:         "teams cue id",
			line:         "5b5f513f-9248-4de5-9744-4147104ec246/12-0",
			hasSpeaker:   false,
			expectedLine: "",
		},
		{
			name:         "teams voice tag new speaker",
			line:         "<v Eric Jorgensen>Previously on a bi-weekly cadence was</v>",
			hasSpeaker:   true,
			expectedLine: " Previously on a bi-weekly cadence was",
		},
		{
			name:           "teams voice tag same speaker",
			line:           "<v Eric Jorgensen>what should we keep doing,</v>",
			currentSpeaker: "Eric Jorgensen",
			hasSpeaker:     false,
			expectedLine:   "what should we keep doing,",
		},
		{
			name:           "teams voice tag without closing tag",
			line:           "<v Eric Jorgensen>Surface blockers ask questions of",
			currentSpeaker: "",
			hasSpeaker:     true,
			expectedLine:   " Surface blockers ask questions of",
		},
		{
			name:           "teams continuation with closing tag",
			line:           "that it was a really helpful checkpoint</v>",
			currentSpeaker: "Eric Jorgensen",
			hasSpeaker:     false,
			expectedLine:   "that it was a really helpful checkpoint",
		},
		{
			name:           "teams continuation without closing tag",
			line:           "directors.",
			currentSpeaker: "Eric Jorgensen",
			hasSpeaker:     false,
			expectedLine:   "directors.",
		},
		{
			name:         "teams voice tag empty text",
			line:         "<v Eric Jorgensen></v>",
			hasSpeaker:   false,
			expectedLine: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			speaker, result := HandleLine(tt.currentSpeaker, tt.line)
			if (speaker != "" && !tt.hasSpeaker) || (speaker == "" && tt.hasSpeaker) {
				t.Errorf("expected hasSpeaker=%t, but got speaker=%q", tt.hasSpeaker, speaker)
			}
			if result != tt.expectedLine {
				t.Errorf("expected %q, but got %q", tt.expectedLine, result)
			}
		})
	}
}
