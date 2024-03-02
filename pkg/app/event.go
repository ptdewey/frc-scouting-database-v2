package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ptdewey/frc-scouting-database-v2/internal/analysis"
	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
)

// TODO: docs
func AnalyzeEvent(eventKey string, apiKey string) ([][]string, error) {
    // Fetch raw match data from api
    rm, err := api.EventMatchesList(eventKey, apiKey)
    if err != nil {
        return nil, err
    }

    // Format match list
    fm, err := api.FormatMatchList(rm)
    if err != nil {
        return nil, err
    }

    // Get raw team list for event
    rt, err := api.TeamList(eventKey, apiKey)
    if err != nil {
        return nil, err
    }

    // Format team list for event
    ft, err := api.FormatTeamList(rt)
    if err != nil {
        return nil, err
    }

    // Reformat output path to include directories for the year and event
    // (Create these directories if necessary)
    outputPath := filepath.Join("output", string([]rune(eventKey)[0:4]), eventKey)
    if err := os.MkdirAll(outputPath, 0755); err != nil {
        return nil, err
    }

    // Extract match statistics from match data and write to CSV
    fname := filepath.Join(outputPath, eventKey + "_matches.csv")
    _, err = analysis.MatchesToCSV(fm, fname)
    if err != nil {
        return nil, err
    }

    // Extract match data for each team in event
    // TODO: parallelize this loop
    for _, t := range ft {
        // Extract team data from overall match data
        tm, err := analysis.TeamMatches(fm, t.Key)
        if err != nil {
            fmt.Println("An error occurred extrating team data for team:", t.Key, err)
            // TEST: potential point of failure here with no matches played for a team
            continue
            // return nil, err
        }

        // Define filename for team data
        fname = filepath.Join(outputPath, eventKey + "_" + t.Key + ".csv")

        // Write team data to csv
        _, err = analysis.TeamMatchesToCSV(tm, t.Key, fname)
        if err != nil {
            fmt.Println("An error occurred writing team data to CSV:", err)
            return nil, err
        }
    }

    // Calculate OPR for event
    opr, err := analysis.CalcEventOPR(fm)
    if err != nil {
        fmt.Println("An error occurred calculating OPRs:", err)
        return nil, err
    }

    // Write OPR data to csv
    fname = filepath.Join(outputPath, eventKey + "_opr.csv")
    fmt.Println(fname)
    out, err := analysis.OPRToCSV(opr, fname)
    if err != nil {
        fmt.Println("An error occurred writing opr to CSV:", err)
        return nil, err
    }

    return out, nil
}
