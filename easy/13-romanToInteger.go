func romanToInt(s string) int {
    romanMap := map[string]int { 
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    
    total := 0
    for i := 0; i < len(s); i++ {
        r1 := romanMap[string(s[i])]
        r2 := 0
        
        // If there are more numerals to the right still
        if i+1 < len(s) {
            r2 = romanMap[string(s[i+1])]
        }
        
        if r1 < r2 {
            romanVal := r2 - r1
            total = total + romanVal
            
            // Advance so that we don't look at
            // r2 again in the next iteration
            i++
        } else {
            total = total + r1
        }
    }
    
    return total
}
