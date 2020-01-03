func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    
    for i := 0; i <= len(haystack)-len(needle); i++ {
        if match(haystack[i:i+len(needle)], needle) {
            return i
        }
    }
    
    return -1
}

func match(s1 string, s2 string) bool {
    if len(s1) == 0 && len(s2) == 0 {
        return true
    }
        
    if s1[0] == s2[0] {
        return match(s1[1:], s2[1:])
    } else {
        return false
    }
}
