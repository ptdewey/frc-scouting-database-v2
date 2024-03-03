package analysis

import (
	"errors"
	"fmt"

	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
	"github.com/ptdewey/frc-scouting-database-v2/internal/utils"
)


// TODO: docs
type AllianceInfo struct {
    Alliances [][]string
    Scores []float64
    AutoScores []float64
    TeleScores []float64
    RP []float64
}


// TODO: docs
func GetEventAlliances(matches []api.Match) (AllianceInfo, []string, error) {
    // Initialize container variables
    var out AllianceInfo

    // Exit if 0 matches were played
    if len(matches) == 0 {
        fmt.Println("Match schedule does not exist.")
        return AllianceInfo{}, nil, errors.New("Match schedule not found for event. OPR cannot be calculated.")
    }
    
    // Iterate through matches populating alliances and scores.
    for _, m := range matches {
        // Skip matches with invalid score breakdown
        if m.ScoreBreakdown == nil {
            continue
        }

        // Remove unplayed matches and playoff matches
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


// TODO: docs
func extractTeamAlliances(alliances AllianceInfo, teamKey string) ([]int, error) {
    var out []int
    for i, a := range alliances.Alliances {
        if utils.ContainsString(teamKey, a) {
            out = append(out, i)
        }
    }
    
    return out, nil
}
