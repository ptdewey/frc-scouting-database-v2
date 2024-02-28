package api

import (
    "encoding/json"
    "log"
)


// Team list item type.
type Team struct {
    Key string `json:"key"`
    Number int `json:"team_number"`
    Name string `json:"nickname"`
}


// Event list item type.
type Event struct {
    Key string `json:"key"`
    Name string `json:"name"`
    EventType int `json:"event_type"`
    StartDate string `json:"start_date"`
    EndDate string `json:"end_date"`
}


// Event match item type
type Match struct {
    Key string `json:"key"`
    MatchNumber int `json:"match_number"`
    CompLevel string `json:"comp_level"`
    Alliances MatchAlliances `json:"alliances"`
    ScoreBreakdown *ScoreBreakdowns `json:"score_breakdown"`
    WinningAlliance string `json:"winning_alliance"`
}


type MatchAlliances struct {
    Blue Alliance `json:"blue"`
    Red Alliance `json:"red"`
}


type Alliance struct {
    TeamKeys [3]string `json:"team_keys"`
    SurrogateTeamKeys []string `json:"surrogate_team_keys"`
    DQTeamKeys []string `json:"dq_team_keys"`
}


type ScoreBreakdowns struct {
    Blue ScoreDetails `json:"blue"`
    Red ScoreDetails `json:"red"`
}


type ScoreDetails struct {
    TotalPoints int `json:"totalPoints"`
    AutoPoints int `json:"autoPoints"`
    TeleopPoints int `json:"teleopPoints"`
    AdjustPoints int `json:"adjustPoints"`
    RP int `json:"rp"`
    // TODO: potentially use map interface to allow for extension to year-specific information
}


// FormatEventList formats the output from api.EventList into a data structure
// containing the desired fields from the original json byte array.
// It takes in a byte array containing the raw json data.
// It outputs an array of type []Event.
// It has no effect on the global state and has no side effects.
func FormatEventList(rawEvents []byte) ([]Event, error) {
    var events []Event
    err := json.Unmarshal(rawEvents, &events)
    if err != nil {
        log.Fatalf("Error conveting from JSON: %v", err)
        return nil, err
    }

    return events, nil
}


// FormatMatchList formats the output from api.EventMatchesList into a data structure
// containing the desired fields from the original json byte array.
// It takes in a byte array containing the raw json data.
// It outputs an array of type []Match.
// It has no effect on the global state and has no side effects.
func FormatMatchList(rawMatches []byte) ([]Match, error) {
    var matches []Match
    err := json.Unmarshal(rawMatches, &matches)
    if err != nil {
        log.Fatalf("Error conveting from JSON: %v", err)
        return nil, err
    }

    return matches, nil
}


// FormatTeamList formats the output from api.TeamList into a data structure
// containing the desired fields from the original json byte array.
// It takes in a byte array containing the raw json data.
// It outputs an array of type Team.
// It has no effect on the global state and has no side effects.
func FormatTeamList(rawTeams []byte) ([]Team, error) {
    var teams []Team
    err := json.Unmarshal(rawTeams, &teams)
    if err != nil {
        log.Fatalf("Error conveting from JSON: %v", err)
        return nil, err
    }

    return teams, nil
}
