package app

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ptdewey/frc-scouting-database-v2/internal/analysis"
	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
)

// Function AnalyzeEvent provides a forward-facing wrapper around the
// backend data fetching and analysis functions.
// It takes in an event key and api key, pulls data from the desired
// event, cleans and formats the data, then calculates summary statistics
// and OPR values for each team participating at the event.
// The final results are written to csv files in the output directory,
// and a 2D string array of the final information is returned as well.
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

    // Extract alliances and scores for each match, for calculating opr and summary stats
    a, teams, err := analysis.GetEventAlliances(fm)
    if err != nil {
        return nil, err
    }

    var stats []analysis.SummaryStats

    // Extract match data for each team in event
    for _, t := range ft {
        // Extract team data from overall match data
        tm, err := analysis.TeamMatches(fm, t.Key)
        if err != nil {
            fmt.Println("An error occurred extrating team data for team:", t.Key, err)
            continue
        }

        // Define filename for team data
        fname = filepath.Join(outputPath, eventKey + "_" + t.Key + ".csv")

        // Write team data to csv
        _, err = analysis.TeamMatchesToCSV(tm, t.Key, fname)
        if err != nil {
            fmt.Println("An error occurred writing team data to CSV:", err)
            return nil, err
        }

        // Calculate team summary statistics
        s, err := analysis.CalcTeamSummaryStats(a, t.Key)
        if err != nil {
            fmt.Println("An error occurred calculating summary stats for team:", t.Key, err)
            return nil, err
        }
        stats = append(stats, s)
    }

    // Write summary stats file
    // TODO: sort first?
    fname = filepath.Join(outputPath, eventKey + "_stats.csv")
    _, err = analysis.SummaryStatsToCSV(stats, fname)
    if err != nil {
        fmt.Println("An error occurred writing summary stats:", err)
        return nil, err
    }
    
    // Calculate OPR for event
    opr, err := analysis.CalcEventOPR(fm, a, teams)
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


// Function AnalyzeActiveEvents pulls TBA api data for all events happening between
// the specified start and end dates, and calls AnalyzeEvent on them.
// This results in the ouput directory being populated with cleaned data and event
// stats for each event in the time range.
// Takes in two strings for the start and end date, and a third string containing
// a valid TBA api key. Can return an error.
func AnalyzeActiveEvents(startDate string, endDate string, apiKey string) error {
    year := string([]rune(startDate)[0:4])

    // Get event list
    re, err := api.EventList(year, apiKey)
    if err != nil {
        return err
    }

    // Format events list
    events, err := api.FormatEventList(re)
    if err != nil {
        return err
    }

    // Reformat time strings to date type
    timeFmt := "2006-01-02"
    start, err := time.Parse(timeFmt, startDate)
    if err != nil {
        return err
    }
    end, err := time.Parse(timeFmt, endDate)
    if err != nil {
        return err
    }

    // Iterate through events, checking if event is considered 'active'
    for _, e := range events {
        // Convert event start and end dates to time type
        evStart, err := time.Parse(timeFmt, e.StartDate)
        if err != nil {
            return err
        }
        evEnd, err := time.Parse(timeFmt, e.EndDate)
        if err != nil {
            return err
        }
        
        // Compare start and end dates
        if (start.After(evStart) || start.Equal(evStart)) && (end.Before(evEnd) || end.Equal(evEnd)) {
            AnalyzeEvent(e.Key, apiKey)
        }
    }

    return nil
}


// Function AnalyzeActiveEvents pulls TBA api data for all events happening 
// currently, and calls AnalyzeEvent on them to extract insights.
// Takes in a string containing the TBA api key and can return an error.
func AnalyzeCurrentEvents(apiKey string) error {
    t := time.Now().Format("2006-01-02")
    err := AnalyzeActiveEvents(t, t, apiKey)
    if err != nil {
        return err
    }

    return nil
}
