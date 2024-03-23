package aggregation

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ptdewey/frc-scouting-database-v2/internal/analysis"
)

type SeasonOPR struct {
	TeamKey     string
	OPRs        []float64
	AutoOPRs    []float64
	TeleopOPRs  []float64
	RPOPRs      []float64
}

// Function CombineEventOPRs takes in a season's worth of OPR data,
// and outputs a map of strings (team keys) to SeasonOPR objects containing
// all opr data for each team in a given season.
func CombineEventOPRs(oprs []analysis.OPR) map[string]*SeasonOPR {
    // TODO: might need to check if event key is not included already to avoid rewriting every time
    // - could maybe pass in opr map mapped to event key to make this easier
    // - this could also potentially also allow for keeping track of event weeks
    // TODO: might also need to take in map as param to allow updating it
    seasonOPRs := make(map[string]*SeasonOPR)
    // iterate over all opr event data
    for _, o := range oprs {
        // update map with teamkey if it does not exist
        if _, exists := seasonOPRs[o.TeamKey]; !exists {
            seasonOPRs[o.TeamKey] = &SeasonOPR{ TeamKey: o.TeamKey }
        }

        // append event stats to map slices
        td := seasonOPRs[o.TeamKey]
        td.OPRs = append(td.OPRs, o.OPR)
        td.AutoOPRs = append(td.AutoOPRs, o.AutoOPR)
        td.TeleopOPRs = append(td.TeleopOPRs, o.TeleopOPR)
        td.RPOPRs = append(td.RPOPRs, o.RPOPR)
    }
    return seasonOPRs
}


// Function calcMaxEvents calculates the max number of events played by a
// team over the course of a given season.
// It takes in a string to SeasonOPR map and returns an integer.
func calcMaxEvents(seasonOPRs map[string]*SeasonOPR) int {
    max := 0
    for _, v := range seasonOPRs {
        // TODO: to avoid checking entire map, stop early over 7?
        if len(v.OPRs) > max {
            max = len(v.OPRs)
        }
    }
    return max 
}


// TODO: docs
// TODO: output to CSV
func SeasonOPRtoCSV(seasonOPRs map[string]*SeasonOPR, year string) error {
    fp := filepath.Join("output", year, year + "_opr.csv")
    f, err := os.Create(fp)
    if err != nil {
        return err
    }
    defer f.Close()

    w := csv.NewWriter(f)
    defer w.Flush()

    maxEvents := calcMaxEvents(seasonOPRs)

    // create and write csv header
    header := []string{ "team_key", "max_opr", "max_auto_opr", "max_tele_opr", "max_rp_opr" }
    for i := 1; i <= maxEvents; i++ {
        header = append( header, "opr_" + strconv.Itoa(i), "auto_opr_" + strconv.Itoa(i),
            "tele_opr_" + strconv.Itoa(i) + "rp_opr_" + strconv.Itoa(i)) 
    }
    w.Write(header)

    // iterate through map of team season data
    for _, ts := range seasonOPRs {
        row := []string{ ts.TeamKey }
        // calculate maximums
        mo, ma, mt, mr := -100.0, -100.0, -100.0, -100.0
        for i := 0; i < maxEvents; i++ {
            mo = max(ts.OPRs[i], mo)
            ma = max(ts.OPRs[i], ma)
            mt = max(ts.OPRs[i], mt)
            mr = max(ts.OPRs[i], mr)
        }
        row = append(row,  
            strconv.FormatFloat(mo, 'f', -1, 64),
            strconv.FormatFloat(ma, 'f', -1, 64), 
            strconv.FormatFloat(mt, 'f', -1, 64),
            strconv.FormatFloat(mr, 'f', -1, 64))

        // add all opr values to row
        for i := 0; i < maxEvents; i++ {
            if i < len(ts.OPRs) {
                row = append(row, 
                    strconv.FormatFloat(ts.OPRs[i], 'f', -1, 64),
                    strconv.FormatFloat(ts.AutoOPRs[i], 'f', -1, 64),
                    strconv.FormatFloat(ts.TeleopOPRs[i], 'f', -1, 64),
                    strconv.FormatFloat(ts.RPOPRs[i], 'f', -1, 64))
            } else {
                row = append(row, "", "", "", "")
                break // TODO: this is ok right?
            }
        }
        w.Write(row)
    }

    return nil
}
