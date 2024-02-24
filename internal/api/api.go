package api

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)


type Team struct {
    Key string `json:"key"`
    Number int `json:"team_number"`
    Name string `json:"name"`
}


// EventList fetches a list of events for a given year from The Blue Alliance API.
// It takes in a string for the year and a Blue Alliance API key.
// The output is a tuple containing a string map interface and an error.
// This function does not modify any global state and has no side effects.
func EventList(year string, apiKey string) ([]map[string]interface{}, error) {
    url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/events/%s/simple", year)

    // make request to tba api and populate events object
    var events []map[string]interface{}
    err := tbaRequest(url, apiKey, events)
    if err != nil {
        log.Fatalf("Error fetching events list: %v", err)
        return nil, err
    }

    return events, nil
}


// EventMatches fetches a json object containing information on all matches for an event.
// It takes in a string for the event key (i.e. "2023vagle") and a Blue Alliance API key.
// The output is a tuple containing a string map interface and an error.
// This function does not modify any global state and has no side effects.
func EventMatches(eventKey string, apiKey string) ([]map[string]interface{}, error){
    url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/event/%s/matches", eventKey)

    // make request to tba api and populate matches object
    var matches []map[string]interface{}
    err := tbaRequest(url, apiKey, matches)
    if err != nil {
        log.Fatalf("Error fetching matches for event %s: %v", eventKey, err)
        return nil, err
    }

    return matches, nil
}


// EventList fetches a list of teams attending a given event from The Blue Alliance API.
// It takes in a string for the event key (i.e. "2023vagle") and a Blue Alliance API key.
// The output is a tuple containing a slice of type Team and an error.
// This function does not modify any global state and has no side effects.
func TeamList(eventKey string, apiKey string) ([]Team, error) {
    url := fmt.Sprintf("https://www.thebluealliance.com/api/v3/event/%s/teams", eventKey)

    // make request to tba api and populate teams object
    var teams []Team
    err := tbaRequest(url, apiKey, teams)
    if err != nil {
        log.Fatalf("Error fetching teams list for event %s: %v", eventKey, err)
        return nil, err
    }
    
    return teams, nil
}


// tbaRequest is a utility function for making requests to the Blue Alliance API.
// It takes in a 2 strings and a pointer to an interface, which is modified in place.
// This function does not modify the global state. 
// This function has the side effect of modifying the object passed in as 'target'.
func tbaRequest(url string, apiKey string, target interface{}) error {
    // create new request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatalf("Error creating request: %v", err)
        return err
    }

    // set required headers
    req.Header.Set("Accept", "application/json")
    req.Header.Set("X-TBA-Auth-Key", apiKey)

    // execute request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // read and parse the json response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %v", err)
        return err
    }

    // unmarshal output into go data structure
    err = json.Unmarshal(body, target)
    if err != nil {
        log.Fatalf("Error unmarshalling body: %v", err)
        return err
    }

    return nil
}

