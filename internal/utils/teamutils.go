package utils

import (
	"errors"

	"github.com/ptdewey/frc-scouting-database-v2/internal/api"
)

// Function to get a subset of matches from an event containing only matches a
// specified team was a part of.
// Takes in a pointer to the event matches object and a team key to search for.
// Returns a slice of matches and error.
// This function does not modify the global state and has no side effects
func TeamMatches(matches *[]api.Match, teamKey string) ([]api.Match, error) {
    // TODO: probably more memory efficient to just return indices of matches
    // where match numbers appear
    var out []api.Match
    for _, m := range *matches {
        if ContainsString(teamKey, m.Alliances.Blue.TeamKeys[:]) {
            out = append(out, m)
        } else if ContainsString(teamKey, m.Alliances.Red.TeamKeys[:]) {
            out = append(out, m)
        }
    }

    // team name not found in any matches
    if len(out) == 0 {
        return nil, errors.New("Team key not found in event matches.")
    }

    return out, nil
}


// Function that writes the output of TeamMatches to a csv file.
// Takes in a pointer to the event matches object and a team key to search for.
// Returns an error if the write failed.
func TeamMatchesToCSV(matches *[]api.Match, teamKey string, filename string) ([][]string, error) {
    // extract team matches from event matches
    tm, err := TeamMatches(matches, teamKey)
    if err != nil {
        return nil, err
    }

    // convert matches struct to csv and write to file
    out, err := MatchesToCSV(tm, filename)
    if err != nil {
        return nil, err
    }
    
    return out, nil 
}
