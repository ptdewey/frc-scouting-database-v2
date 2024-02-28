package utils


// Function to check if a string is in a slice of strings.
// Takes in a string and slice of strings, returns t/f if string is within the slice.
func ContainsString(str string, slice []string) bool {
    for _, s := range slice {
        if str == s {
            return true
        }
    }
    return false
}
