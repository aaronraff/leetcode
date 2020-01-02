var mappings map[int]string = map[int]string {
    2: "abc",
    3: "def",
    4: "ghi",
    5: "jkl",
    6: "mno",
    7: "pqrs",
    8: "tuv",
    9: "wxyz",
}

func letterCombinations(digits string) []string {
    if len(digits) < 1 {
        return []string{}
    }
    
    combs := make([]string, 0)
    firstMapping := mappings[int(digits[0] - '0')]

    for _, char := range(firstMapping) {
        combs = append(combs, string(char))
    }
    
    // Remove the first digit since we just added that
    digits = digits[1:]

    for _, digit := range(digits) {
        mapping := mappings[int(digit - '0')]
        removeUpTo := len(combs) - 1
        
        for _, comb := range(combs) {
            for _, char := range(mapping) {
                combs = append(combs, comb + string(char))
            }
        }

        // Remove what we had before adding onto the strings
        combs = combs[removeUpTo+1:]
    }
    
    return combs
}
