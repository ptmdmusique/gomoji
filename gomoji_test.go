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

func hasEmoji(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
func TestGetEmoji(t *testing.T) {
	emoji := GetEmoji(Positive)
	if !hasEmoji(PosEmoList, emoji) {
		t.Errorf("Expect negative emoji, got positive emoji")
	}

	emoji = GetEmoji(Negative)
	if !hasEmoji(NegEmoList, emoji) {
		t.Errorf("Expect positive emoji, got negative emoji")
	}
}
