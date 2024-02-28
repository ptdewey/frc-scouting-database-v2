package utils

import "testing"


// Test ContainsString function from utils package
func TestContainsString(t *testing.T) {
    slice := []string{"hello", "world"}
    if !ContainsString("hello", slice) {
        t.Fatalf("ContainsString missed member string.")
    }
    if !ContainsString("world", slice) {
        t.Fatalf("ContainsString missed member string.")
    }
    if ContainsString("text", slice) {
        t.Fatalf("ContainsString incorrectly identified non-member string.")
    }
}


// Test ExtractUniqueStrings for correct output
func TestExtractUniqueStrings(t *testing.T) {
    arr := [][]string {
        {"hello", "world", "!"},
        {"test", "hello", "world"},
        {"test", "world", "hello"},
        {"foo", "bar", "hello"},
    }

    // expected output
    u := []string {"hello", "world", "!", "test", "foo", "bar"}

    res := ExtractUniqueStrings(arr)

    // check for valid lengths
    if len(u) != len(res) {
        t.Fatalf("Mismatched output lengths. %d %d", len(u), len(res)) 
    }
    
    // check unique strings were correctly extracted
    for i, s := range u {
        // output order can vary so check if each string is contained
        if !ContainsString(s, res) {
            t.Fatalf("Unexpected output at index %d. %s %s", i, s, res[i])
        }
    }
}
