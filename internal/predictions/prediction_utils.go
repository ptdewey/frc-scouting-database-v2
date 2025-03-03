package predictions

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ptdewey/frc-scouting-database-v2/internal/aggregation"
	"github.com/ptdewey/frc-scouting-database-v2/internal/analysis"
)

type Alliance struct {
	TeamKeys   []string
	OPRs       []float64
	AutoOPRs   []float64 // NOTE: this can be used to cap scores to make better predictions
	TeleopOPRs []float64
	RPOPRs     []float64
}

// TODO: docs
// TODO: get predicted contribution values for alliance
func getAllianceOPR(teamKeys []string, season map[string]*aggregation.SeasonOPR) (Alliance, error) {
	var a Alliance

	// populate alliance object
	for _, t := range teamKeys {
		o, err := getTeamOPR(t, season)
		if err != nil {
			return Alliance{}, err
		}
		a.TeamKeys = append(a.TeamKeys, t)
		a.OPRs = append(a.OPRs, o.OPR)
		a.AutoOPRs = append(a.AutoOPRs, o.AutoOPR)
		a.TeleopOPRs = append(a.TeleopOPRs, o.TeleopOPR)
		a.RPOPRs = append(a.RPOPRs, o.RPOPR)
	}

	return a, nil
}

// TODO: docs
// TODO: Get team opr (max/weighted mean/adjusted max) from map
func getTeamOPR(teamKey string, season map[string]*aggregation.SeasonOPR) (analysis.OPR, error) {
	var out analysis.OPR

	// get team season data from map
	s, ok := season[teamKey]
	if !ok {
		return analysis.OPR{}, errors.New("Team key does not exist.")
	}

	// populate opr object with best contribution ratings
	for i := range len(s.OPRs) {
		out.OPR = max(out.OPR, s.OPRs[i])
		out.AutoOPR = max(out.AutoOPR, s.AutoOPRs[i])
		out.TeleopOPR = max(out.TeleopOPR, s.TeleopOPRs[i])
		out.RPOPR = max(out.RPOPR, s.RPOPRs[i])
	}

	return out, nil
}

// TODO: docs
// TODO: Calculate the average contribution ratings for a season
// NOTE: this is for schedule strength calculations
func calcAverageContribution() (analysis.OPR, error) {
	return analysis.OPR{}, nil
}

// DOC:
func WritePredMatchesToCSV(fname string, matches []PredMatch) error {
	file, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// TODO: Add match numbers (they seem out of order though)
	header := []string{"BlueAlliance", "BlueScore", "BlueAutoScore", "BlueTeleopScore", "BlueSumScore", "BlueRP",
		"RedAlliance", "RedScore", "RedAutoScore", "RedTeleopScore", "RedSumScore", "RedRP",
		"WinningAlliance", "WinningMargin"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	for _, match := range matches {
		record := []string{
			// fmt.Sprintf("%d", i+1),
			strings.Join(match.BlueScores.TeamKeys, ";"),
			fmt.Sprintf("%.2f", match.BlueScores.Score),
			fmt.Sprintf("%.2f", match.BlueScores.AutoScore),
			fmt.Sprintf("%.2f", match.BlueScores.TeleopScore),
			fmt.Sprintf("%.2f", match.BlueScores.SumScore),
			fmt.Sprintf("%.2f", match.BlueScores.RP),
			strings.Join(match.RedScores.TeamKeys, ";"),
			fmt.Sprintf("%.2f", match.RedScores.Score),
			fmt.Sprintf("%.2f", match.RedScores.AutoScore),
			fmt.Sprintf("%.2f", match.RedScores.TeleopScore),
			fmt.Sprintf("%.2f", match.RedScores.SumScore),
			fmt.Sprintf("%.2f", match.RedScores.RP),
			match.WinningAlliance,
			fmt.Sprintf("%.2f", match.WinningMargin),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}

	return nil
}
