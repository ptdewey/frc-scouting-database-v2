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
