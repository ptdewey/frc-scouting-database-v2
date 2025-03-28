package analysis

import (
	"errors"
	"fmt"

	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
	"github.com/ptdewey/frc-scouting-database-v2/internal/utils"
)

type AllianceScoreBreakdown struct {
	Alliances  [][]string
	Scores     []float64
	AutoScores []float64
	TeleScores []float64
	RP         []float64
}

// Function GetEventAllianceScores is a helper function for statistical analysis functions
// that extracts team and score info from match breakdowns.
// It takes in a api.Match slice and outputs an AllianceInfo object, a string slice
// containing all unique team keys for an event, and an error.
func GetEventAllianceScores(matches []api.Match) (AllianceScoreBreakdown, []string, error) {
	// Initialize container variables
	var out AllianceScoreBreakdown

	// Exit if 0 matches were played
	if len(matches) == 0 {
		fmt.Println("Match schedule does not exist.")
		return AllianceScoreBreakdown{}, nil, errors.New("Match schedule not found for event. OPR cannot be calculated.")
	}

	// Iterate through matches populating alliances and scores.
	for _, m := range matches {
		// Skip matches with invalid score breakdown
		if m.ScoreBreakdown == nil {
			continue
		}

		// Remove unplayed matches and playoff matches
		// TODO: allow sf, f matches in calculation?
		if m.ScoreBreakdown.Blue.TotalPoints == -1 || m.CompLevel != "qm" {
			continue
		}

		// Append metrics to output struct
		out.Alliances = append(out.Alliances, m.Alliances.Blue.TeamKeys[:])
		out.Alliances = append(out.Alliances, m.Alliances.Red.TeamKeys[:])
		out.Scores = append(out.Scores, float64(m.ScoreBreakdown.Blue.TotalPoints))
		out.Scores = append(out.Scores, float64(m.ScoreBreakdown.Red.TotalPoints))
		out.AutoScores = append(out.AutoScores, float64(m.ScoreBreakdown.Blue.AutoPoints))
		out.AutoScores = append(out.AutoScores, float64(m.ScoreBreakdown.Red.AutoPoints))
		out.TeleScores = append(out.TeleScores, float64(m.ScoreBreakdown.Blue.TeleopPoints))
		out.TeleScores = append(out.TeleScores, float64(m.ScoreBreakdown.Red.TeleopPoints))
		out.RP = append(out.RP, float64(m.ScoreBreakdown.Blue.RP))
		out.RP = append(out.RP, float64(m.ScoreBreakdown.Red.RP))
	}

	// Extract unique team keys.
	teams := utils.ExtractUniqueStrings(out.Alliances)

	return out, teams, nil
}

// Function extractTeamAlliances is a helper function that extracts the indices
// within the AllianceInfo table in which a team key occurs (team plays in a match).
// It takes in an AllianceInfo object and a string teamKey.
// It ouputs an integer slice and an error.
func extractTeamAlliances(alliances AllianceScoreBreakdown, teamKey string) ([]int, error) {
	var out []int
	for i, a := range alliances.Alliances {
		if utils.ContainsString(teamKey, a) {
			out = append(out, i)
		}
	}

	return out, nil
}
