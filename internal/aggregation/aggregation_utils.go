package aggregation

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ptdewey/frc-scouting-database-v2/internal/analysis"
)

// TODO: it would be more efficient to just take in an array of opr objects each time new data is fetched

// TODO: docs
func combineOPRFiles(year string) ([]analysis.OPR, error) {
    // TODO: iterate through each eventKey in output/year directory

    return nil, nil
}


// TODO: docs
func readEventOPR(eventKey string, year string) ([]analysis.OPR, error) {
    var out []analysis.OPR

    // read opr file from eventKey directory
    fp := filepath.Join("output", year, eventKey, eventKey + "_opr.csv")
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

        opr, err := csvToOPR(row)
        if err != nil {
            return nil, err
        }

        // append converted row to output structure
        out = append(out, opr)
    }

    return out, nil 
}


// TODO: docs
func csvToOPR(row []string) (analysis.OPR, error) {
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
