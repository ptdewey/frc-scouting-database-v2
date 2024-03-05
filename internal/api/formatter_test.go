package api

import "testing"

// TestEventList calls api.EventList and api.FormatEventList checking for correct output.
func TestFormatEventList(t *testing.T) {
    year := "2024"
    out, err := EventList(year, apiKey)
    if err != nil {
        t.Fatalf(`EventList(year, apiKey) = %q. Error: %v`, out, err)
    }

    // run formatter function
    events, err := FormatEventList(out)
    if err != nil {
        t.Fatalf(`FormatEventList(rawEvents) = %q. Error: %v`, out, err)
    }
    
    // event index in output array
    ei := 171

    // test correct event key
    ek := "2024vagle"
    if events[ei].Key != ek {
        t.Fatalf("Event at index %d did not match key %s %s.", ei, ek, events[ei].Key)
    }

    // test for correct event name
    en := "CHS District Glen Allen VA Event"
    if events[ei].Name != en {
        t.Fatalf("Event at index %d had mismatched event names %s %s.", ei, en, events[ei].Name)
    }

    // test for correct event type
    et := uint8(1)
    if events[ei].EventType != et {
        t.Fatalf("Event at index %d had mismatched event type %d %d.", ei, et, events[ei].EventType)
    }

    // test for correct start date
    sd := "2024-03-15"
    if events[ei].StartDate != sd {
        t.Fatalf("Event at index %d had mismatched start date %s %s.", ei, sd, events[ei].StartDate)
    }

    // test for correct end date
    ed := "2024-03-17"
    if events[ei].EndDate != ed {
        t.Fatalf("Event at index %d had mismatched end date %s %s.", ei, ed, events[ei].EndDate)
    }
}


// TestEventMatchesList calls api.EventMatchesList and api.FormatMatchList checking for correct output.
func TestFormatEventMatchesList(t *testing.T) {
    eventKey := "2020vagle"
    out, err := EventMatchesList(eventKey, apiKey)
    if err != nil {
        t.Fatalf(`EventMatches(eventKey, apiKey) = %q. Error: %v`, out, err)
    }

    // test formatting function
    matches, err := FormatMatchList(out)
    if err != nil {
        t.Fatalf(`FormatMatchList(rawMatches) = %q. Error %v`, out, err)
    }
    
    // index of match in output array
    mi := 5

    // test for correct match key
    mk := "2020vagle_qf1m3"
    if matches[mi].Key != mk {
        t.Fatalf("Match at index %d did not match key %s.", mi, mk)
    }

    // test for correct match number
    mn := uint8(3)
    if matches[mi].MatchNumber != mn {
        t.Fatalf("Match at index %d did not match match number %d.", mi, mn)
    }

    // test for correct comp level
    cl := "qf"
    if matches[mi].CompLevel != cl {
        t.Fatalf("Match at index %d has mismatched comp levels %s %s.", 
            mi, cl, matches[mi].CompLevel)
    }

    // test for correct match alliances
    ra := [3]string{"frc1610", "frc2106", "frc1123"}
    if matches[mi].Alliances.Red.TeamKeys != ra {
        t.Fatalf("Expected red alliance %s at %d did not match red alliance %s.",
            ra, mi, matches[mi].Alliances.Red.TeamKeys)
    }
    ba := [3]string{"frc2998", "frc401", "frc6334"}
    if matches[mi].Alliances.Blue.TeamKeys != ba {
        t.Fatalf("Expected blue alliance %s at %d did not match blue alliance %s.",
            ra, mi, matches[mi].Alliances.Blue.TeamKeys)
    }

    // test for correct score breakdown
    if matches[mi].ScoreBreakdown == nil {
        t.Fatalf("Score breakdown at index %d is nil.", mi)
    }
    sdr := ScoreDetails{174, 51, 117, 0, 0}
    if matches[mi].ScoreBreakdown.Red != sdr {
        t.Fatalf("Score breakdown at index %d did not match expected breakdown", mi)
    }
    sdb := ScoreDetails{132, 36, 96, 0, 0}
    if matches[mi].ScoreBreakdown.Blue != sdb {
        t.Fatalf("Score breakdown at index %d did not match expected breakdown", mi)
    }

    // test for correct winning alliance
    wa := "red"
    if matches[mi].WinningAlliance != wa {
        t.Fatalf("Match at index %d had mismatched winning alliances %s %s", 
            mi, wa, matches[mi].WinningAlliance)
    }
}


// TestFormatTeamList calls and api.TeamList api.FormatTeamList checking for correct output
func TestFormatTeamList(t *testing.T) {
    eventKey := "2020vagle"
    out, err := TeamList(eventKey, apiKey)
    if err != nil {
        t.Fatalf(`TeamList(eventKey, apiKey) = %q. Error: %v`, out, err)
    }

    // test formatting function
    teams, err := FormatTeamList(out)
    if err != nil {
        t.Fatalf(`FormatTeamList(rawTeams) = %q. Error: %v`, out, err)
    }

    // index of team within output array
    teamIndex := 13

    teamKey := "frc2106"
    if teams[teamIndex].Key != teamKey {
        t.Fatalf("Team at index %d did not match key %s.", teamIndex, teamKey)
    }

    teamName := "The Junkyard Dogs"
    if teams[teamIndex].Name != teamName {
        t.Fatalf("Team at index %d did not match number %s.", teamIndex, teamName)
    }

    teamNum := uint16(2106)
    if teams[teamIndex].Number != teamNum {
        t.Fatalf("Team at index %d did not match number %d.", teamIndex, teamNum)
    }
}
