package analysis

import (
	// "errors"
	// "fmt"

	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
)


type SummaryStats struct {
    TeamKey string
    MeanScore float64
    MedianScore float64
    MaxScore float64 // TODO: maybe make this a weighted metric to reduce outliers
    MeanAuto float64
}

func CalcSummaryStats(matches []api.Match) (error) {
    // var alliances [][]string
    // var scores, autoScores, teleScores, rp []float64
    //
    // // Exit if 0 matches were played
    // if len(matches) == 0 {
    //     fmt.Println("Match schedule does not exist.")
    //     return nil, errors.New("Match schedule not found for event. OPR cannot be calculated.")
    // }

    // TODO: opr match score table can probably be used here
    // - maybe migrate that part to a utils file
    return nil
}
