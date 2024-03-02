package analysis

import "github.com/ptdewey/frc-scouting-database-v2/internal/api"


type SummaryStats struct {
    TeamKey string
    MeanScore float64
    MedianScore float64
    MaxScore float64 // TODO: maybe make this a weighted metric to reduce outliers
    MeanAuto float64
}

func CalcSummaryStats(matches []api.Match) (error) {

    return nil
}
