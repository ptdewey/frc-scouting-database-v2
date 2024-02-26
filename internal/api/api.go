package api

import (
    "fmt"
    "io"
    "log"
    "net/http"
)


// EventList fetches a list of events for a given year from The Blue Alliance API.
// It takes in a string for the year and a Blue Alliance API key.
// The output is a tuple containing a string map interface and an error.
// This function does not modify any global state and has no side effects.
func EventList(year string, apiKey string) ([]byte, error) {
    url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/events/%s/simple", year)

    // make request to tba api and populate events object
    events, err := tbaRequest(url, apiKey)
    if err != nil {
        log.Fatalf("Error fetching events list: %v", err)
        return nil, err
    }

    return events, nil
}


// EventMatchesList fetches a json object containing information on all matches for an event.
// It takes in a string for the event key (i.e. "2023vagle") and a Blue Alliance API key.
// The output is a tuple containing a string map interface and an error.
// This function does not modify any global state and has no side effects.
func EventMatchesList(eventKey string, apiKey string) ([]byte, error){
    url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/event/%s/matches", eventKey)

    // make request to tba api and populate matches object
    matches, err := tbaRequest(url, apiKey)
    if err != nil {
        log.Fatalf("Error fetching matches for event %s: %v", eventKey, err)
        return nil, err
    }

    return matches, nil
}


// EventList fetches a list of teams attending a given event from The Blue Alliance API.
// It takes in a string for the event key (i.e. "2023vagle") and a Blue Alliance API key.
// The output is a json string
// This function does not modify any global state and has no side effects.
func TeamList(eventKey string, apiKey string) ([]byte, error) {
    url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/event/%s/teams", eventKey)

    // make request to tba api and populate teams object
    teams, err := tbaRequest(url, apiKey)
    if err != nil {
        log.Fatalf("Error fetching teams list for event %s: %v", eventKey, err)
        return nil, err
    }

    return teams, nil
}


// tbaRequest is a utility function for making requests to the Blue Alliance API.
// It takes in a 2 strings and outputs a json string
// This function does not modify the global state. 
func tbaRequest(url string, apiKey string) ([]byte, error) {
    // create new request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalf("Error creating request: %v", err)
        return nil, err
    }

    // set required headers
    req.Header.Set("Accept", "application/json")
    req.Header.Set("X-TBA-Auth-Key", apiKey)

    // execute request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // read and parse the json response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %v", err)
        return nil, err
    }

    return body, nil
}

