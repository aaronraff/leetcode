import "strings"

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    
    rows := make([]string, numRows)    
    strArr := strings.Split(s, "")
    currRow := 0
    goUp := false
    
    for _, char := range(strArr) {
        rows[currRow] = rows[currRow] + char
        
        if goUp {
            currRow--
        } else {
            currRow++
        }
        
        if currRow < 0 {
            goUp = false
            currRow = currRow + 2
        } else if currRow == numRows {
            goUp = true
            currRow = currRow - 2
        }
    }
    
    return strings.Join(rows, "")
}
