import "strings"

func longestPalindrome(s string) string {
    strArr := strings.Split(s, "")
    longest := ""
    var lPos, rPos int
    var match, sameLetter bool
    
    for middle, _ := range(strArr) {
        lPos = middle
        rPos = middle
        match = true
        sameLetter = true
        
        // Start at the middle and go outwards
        for match {
            // Check if index out of bounds
            canMoveLPos := lPos - 1 >= 0
            canMoveRPos := rPos + 1 < len(strArr)
            
            if canMoveLPos && canMoveRPos && strArr[lPos-1] == strArr[rPos+1] {
                lPos--
                rPos++
            } else if canMoveLPos && strArr[lPos-1] == strArr[lPos] && sameLetter {
                lPos--
            } else if canMoveRPos && strArr[rPos+1] == strArr[rPos] && sameLetter {
                rPos++
            } else {
                match = false
            }
            
            // Not a continous string of the same letter
            // i.e "aaaaaa"
            if strArr[lPos] != strArr[middle] || strArr[rPos] != strArr[middle] {
                sameLetter = false
            }
            
            if (rPos - lPos + 1) > len(longest) {
                longest = strings.Join(strArr[lPos:(rPos+1)], "")
            }
        }
    }
                                     
    return longest
}
