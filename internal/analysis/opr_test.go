package analysis

import (
    "fmt"
    "testing"
)


// TestsCalcEventOPR tests for correct output from CalcEventOPR.
func TestCalcEventOPR(t *testing.T) {
    out, err := CalcEventOPR(matches) 
    if err != nil {
        t.Fatalf("%v", err)
    }
    fmt.Println(out)
    // TODO: test output for correctness (expect some floating point variance)
}
