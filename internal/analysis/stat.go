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


// Function CalcTeamSummaryStats calculates summary statistics (mean, max, median) for
// each team at an event. It takes in an AllianceInfo object from alliance_utils.go
// and a string definining which team to evaluate statistics for.
// It returns a SummaryStats object and an error.
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


// Function SummaryStatsToCSV converts an array of SummaryStats to a 2D string
// array and writes the output to a csv file. It takes in a SummaryStats slice
// and a string for the filename and returns a 2D string slice and an error.
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


// Function SummaryStatsToCSVRow converts a SummaryStats object to a csv row string.
// It takes in a SummaryStats object and outputs a string slice.
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
