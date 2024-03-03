package analysis

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
)


type SummaryStats struct {
    TeamKey string
    MeanScore float64
    // MedianScore float64 // TODO: add median later
    MaxScore float64 // TODO: maybe make this a weighted metric to reduce outliers
    MeanAuto float64
    MeanTeleop float64
}


// TODO: docs
func CalcTeamSummaryStats(a AllianceInfo, teamKey string) (SummaryStats, error) {
    // Initialize output variable.
    var out SummaryStats

    // Extract team match indices
    teamMatchIndices, err := extractTeamAlliances(a, teamKey)
    if err != nil {
        return SummaryStats{}, err
    }

    // Iterate through team matches calculating summary statistics.
    n := len(teamMatchIndices)
    for _, i := range teamMatchIndices {
        out.MeanScore = out.MeanScore + a.Scores[i]
        out.MaxScore = math.Max(out.MaxScore, a.Scores[i])
        out.MeanAuto = out.MeanAuto + a.AutoScores[i]
        out.MeanTeleop = out.MeanTeleop + a.TeleScores[i]
    }
    out.TeamKey = teamKey
    out.MeanScore = out.MeanScore / float64(n)
    out.MeanAuto = out.MeanAuto / float64(n)
    out.MeanTeleop = out.MeanTeleop / float64(n)

    return out, nil
}


// TODO: docs
func SummaryStatsToCSV(stats []SummaryStats, filename string) ([][]string, error) {
    // create and open new file
    f, err := os.Create(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    // Create new csv file writer
    w := csv.NewWriter(f)
    defer w.Flush()

    // create 2D array to store rows
    var out [][]string
    
    // csv header
    h := []string {
        "team_key", "mean_score", "max_score", "mean_auto", "mean_tele",
    }
    out = append(out, h)

    // iterate through oprs, adding to out
    for _, s := range stats {
        out = append(out, SummaryStatsToCSVRow(s))
    }
    
    // write out to csv file
    if err := w.WriteAll(out); err != nil {
        return nil, err
    }

    return out, nil
}

// TODO: docs
func SummaryStatsToCSVRow(s SummaryStats) []string {
    var out []string

    // Append opr data to output row
    out = append(out,
        s.TeamKey,
        fmt.Sprint(s.MeanScore),
        fmt.Sprint(s.MaxScore),
        fmt.Sprint(s.MeanAuto),
        fmt.Sprint(s.MeanTeleop),
    )

    return out
}
