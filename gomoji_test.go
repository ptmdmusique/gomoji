package gomoji

import (
	"testing"
)

func TestAnalyzeSentiment(t *testing.T) {
	text := "This is awful"
	got := AnalyzeSentiment(text)
	if got != 0 {
		t.Errorf("Expect negative, got positive")
	}

	text = "This is beautiful"
	got = AnalyzeSentiment(text)
	if got != 1 {
		t.Errorf("Expect positive, got negative")
	}
}
