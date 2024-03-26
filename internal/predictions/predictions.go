package predictions

import (
	"github.com/ptdewey/frc-scouting-database-v2/internal/aggregation"
)

// TODO: decide between score and sumscore as final metric
type PredScore struct {
    TeamKeys []string
    Score float64
    AutoScore float64
    TeleopScore float64
    SumScore float64 
    RP float64
}

type PredMatch struct {
    BlueScores PredScore
    RedScores PredScore
}


// TODO: docs
func PredictAllianceScore(teamKeys []string, season map[string]*aggregation.SeasonOPR) (PredScore, error) {
    var out PredScore 

    // get alliance contribution data
    alliance, err := getAllianceOPR(teamKeys, season)
    if err != nil {
        return PredScore{}, err
    }

    // populate predicted score object
    out.TeamKeys = alliance.TeamKeys
    for i := range len(out.TeamKeys) {
        out.Score = out.Score + alliance.OPRs[i]
        // TODO: more complicated logic for auto score (will overpredict for good alliances)
        out.AutoScore = out.AutoScore + alliance.AutoOPRs[i]
        out.TeleopScore = out.TeleopScore + alliance.TeleopOPRs[i]
        out.RP = out.RP + alliance.RPOPRs[i]
    }
    out.SumScore = out.TeleopScore + out.AutoScore
    
    return out, nil
}


// TODO: docs
func PredictMatchResult(blueTeamKeys []string, redTeamKeys []string, season map[string]*aggregation.SeasonOPR) (PredMatch, error) {
    // calculate blue alliance predicted scores
    bp, err := PredictAllianceScore(blueTeamKeys, season)
    if err != nil {
        return PredMatch{}, err
    }
    // calculate red alliance predicted scores
    rp, err := PredictAllianceScore(blueTeamKeys, season)
    if err != nil {
        return PredMatch{}, err
    }
    
    // return predicted results
    out := PredMatch {
        BlueScores: bp,
        RedScores: rp,
    }

    return out, nil
}
