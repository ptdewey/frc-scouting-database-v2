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


// Function ExtractUniqueStrings find unique strings from 2D string array.
// Takes in a 2D [][]string array and returns a slice of unique strings.
// Output order can vary between runs due to the usage of a map.
func ExtractUniqueStrings(array [][]string) []string {
    uniqueMap := make(map[string]struct{})

    // Iterate through the array
    for _, arr := range array {
        for _, str := range arr {
            // Flag string as found
            uniqueMap[str] = struct{}{}
        }
    }

    // Convert the map keys to a slice of strings
    var unique []string
    for key := range uniqueMap {
        unique = append(unique, key)
    }

    return unique
}
