package app

import (
	"github.com/ptdewey/frc-scouting-database-v2/internal/aggregation"
)

// AggregateEvents aggregates OPR data for all events within a specified year into a single CSV file.
// It serves as a high-level function that orchestrates the process of reading event-specific OPR data from CSV files
// located in the "output" directory (in project root), combining this data into a comprehensive
// map of SeasonOPR objects, and then writing this aggregated data back to a new CSV file. The function ensures that the
// aggregation results in a valid dataset by checking the size of the resulting OPR map; it proceeds to write the results
// only if the map is non-empty. This automated process simplifies the generation of season-wide OPR summaries, but returns
// an error if any step in the read-combine-write pipeline fails or if the aggregated data is deemed invalid due to its length.
func AggregateEvents(year string) (map[string]*aggregation.SeasonOPR, error) {
	outputPath := "output"

	// aggregate season opr files
	season, err := aggregation.ReadAndCombineEventOPRs(outputPath, year)
	if err != nil {
		return nil, err
	}

	// write results to CSV file
	err = aggregation.SeasonOPRtoCSV(season, outputPath, year)
	if err != nil {
		return nil, err
	}

	return season, nil
}
