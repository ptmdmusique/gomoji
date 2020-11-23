package gomoji

import (
	"net/http"

	"github.com/cdipaolo/sentiment"

	"fmt"
	"log"

	"math/rand"
	"time"

	"encoding/json"
)

const (
	Negative = 0
	Positive = 1
)

func AnalyzeSentiment(text string) uint8 {
	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}

	var analysis *sentiment.Analysis
	analysis = model.SentimentAnalysis(text, sentiment.English)

	return analysis.Score
}

func GetEmoji(sentiment uint8) string {
	posList := []string{"🎁", "😙", "💞", "💃", "🎊", "🏆", "☺", "🐾", "😋", "😛", "🌸", "🐱", "😃", "🍜", "💪"}
	negList := []string{"👿", "😕", "😐", "😒", "😿", "😦", "😾", "😠", "👺", "😡", "😨", "💩", "😭", "😓", "👹"}

	rand.Seed(time.Now().Unix())
	var emoji string
	if sentiment == Positive {
		emoji = posList[rand.Int()%len(posList)]
	} else if sentiment == Negative {
		emoji = negList[rand.Int()%len(negList)]
	} else {
		emoji = "You broke the server!"
	}

	return emoji
}

func GetGomoji(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		fmt.Fprint(w, "Invalid request\n")
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sentiment := AnalyzeSentiment(reqBody.Text)
	emoji := GetEmoji(sentiment)
	log.Printf("Got emoji %s", emoji)

	type ResBody struct {
		Emoji     string `json:"emoji"`
		Sentiment uint8  `json:"sentiment"`
	}
	resBody := ResBody{Emoji: emoji, Sentiment: sentiment}

	fmt.Printf("Returning %+v\n", resBody)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resBody)
	return
}
