package aggregation

import (
	"encoding/csv"
    "fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ptdewey/frc-scouting-database-v2/internal/analysis"
)

// combineOPRFiles aggregates OPR data from all event CSV files within a specific year's directory.
// It navigates through each subdirectory of outputPath/year, expecting each to contain an event-specific "_opr.csv" file.
// The function compiles a combined slice of OPR structures from these files, or returns an error if any issues arise during file reading.
// This allows for a consolidated view of OPR data across multiple events for the specified year.
func combineOPRFiles(outputPath string, year string) ([]analysis.OPR, error) {
    var out []analysis.OPR

    // walk through all files in output directory
    outputPath = filepath.Join(outputPath, year)
    err := filepath.WalkDir(outputPath, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }

        // skip non-directory entries
        if !d.IsDir() || d.Name() == year {
            return nil
        }
        
        // call readEventOPR on each file in the directory
        opr, err := readEventOPR(d.Name(), outputPath)
        if err != nil {
            fmt.Println("Error reading contents", d.Name(), ":", err)
            return nil
        }
        out = append(out, opr...)
        
        return err
    })
    if err != nil {
        return nil, err
    }
    
    return out, nil
}


// readEventOPR opens and reads an OPR CSV file for a specific event from the given outputPath.
// It expects the file to be named "<eventKey>_opr.csv" and located within an eventKey subdirectory of outputPath.
// Returns a slice of OPR structures populated with the CSV data, or an error if the file cannot be opened or read.
// Assumes the first row is a header and skips it, parsing subsequent rows into OPR data.
func readEventOPR(eventKey string, outputPath string) ([]analysis.OPR, error) {
    var out []analysis.OPR

    // read opr file from eventKey directory
    fp := filepath.Join(outputPath, eventKey, eventKey + "_opr.csv")
    file, err := os.Open(fp)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // create new csv file reader and read header
    reader := csv.NewReader(file)
    _, err = reader.Read()
    if err != nil {
        return nil, err
    }

    for {
        // read next line
        row, err := reader.Read()
        // check if eof has been reached
        if err != nil {
            break
        }
        opr, err := csvRowToOPR(row)
        if err != nil {
            return nil, err
        }
        // append converted row to output structure
        out = append(out, opr)
    }

    return out, nil 
}


// Function csvRowToOPR converts an opr csv file row string to a opr object.
// It takes in a string array and outputs an analysis.OPR type and error.
func csvRowToOPR(row []string) (analysis.OPR, error) {
    var out analysis.OPR
    // opr
    opr, err := strconv.ParseFloat(row[1], 64)
    if err != nil {
        return out, err
    }
    // auto opr
    apr, err := strconv.ParseFloat(row[2], 64)
    if err != nil {
        return out, err
    }
    // tele opr
    tpr, err := strconv.ParseFloat(row[3], 64)
    if err != nil {
        return out, err
    }
    // rp opr
    rpr, err := strconv.ParseFloat(row[4], 64)
    if err != nil {
        return out, err
    }
    // populate output struct
    out.TeamKey = row[0]
    out.OPR = opr
    out.AutoOPR = apr
    out.TeleopOPR = tpr
    out.RPOPR = rpr
    return out, nil
}
