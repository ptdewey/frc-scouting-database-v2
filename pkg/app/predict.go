package app

import (
	"log"

	"github.com/ptdewey/frc-scouting-database-v2/internal/aggregation"
	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
	"github.com/ptdewey/frc-scouting-database-v2/internal/predictions"
)

// DOCS:
// TEST: test
func PredictCurrentEvent(seasonOPRs map[string]*aggregation.SeasonOPR, eventKey string, apiKey string) ([]predictions.PredMatch, error) {
	// This code is pulled from event.go, there may be a better way to obtain this info (avoiding redundant calls)
	rm, err := api.EventMatchesList(eventKey, apiKey)
	if err != nil {
		return nil, err
	}
	matches, err := api.FormatMatchList(rm)
	if err != nil {
		return nil, err
	}

	preds := make([]predictions.PredMatch, len(matches))

	for i, match := range matches {
		p, err := predictions.PredictMatchResult(match.Alliances.Blue.TeamKeys[:], match.Alliances.Red.TeamKeys[:], seasonOPRs)
		if err != nil {
			log.Println("Failed to generate prediction for match: ", match.Key, match.MatchNumber, err)
		}
		preds[i] = p
	}

	return preds, nil
}
