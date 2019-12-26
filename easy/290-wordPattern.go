import "strings"

func wordPattern(pattern string, str string) bool {
    strArr := strings.Split(str, " ")
    pattArr := strings.Split(pattern, "")
    
    // Need both maps for bijection
    // pattern => str
    map1 := make(map[string]string)
    
    // str => pattern
    map2 := make(map[string]string)
    
    if len(strArr) != len(pattArr) {
        return false
    }
    
    for index, _ := range(pattArr) {
        val1, ok1 := map1[pattArr[index]]
        val2, ok2 := map2[strArr[index]]
        
        if !ok1 {
            map1[pattArr[index]] = strArr[index]
            val1 = strArr[index]
        }
        
        if !ok2 {
            map2[strArr[index]] = pattArr[index]
            val2 = pattArr[index]
        }
        
        // Check bijection
        if val1 != strArr[index] || val2 != pattArr[index] {
            return false
        }
    }
    
    return true
}
