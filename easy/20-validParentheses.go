var matches map[rune]rune = map[rune]rune {
    ')': '(',
    ']': '[',
    '}': '{',
}

func isValid(s string) bool {
    stack := make([]rune, 0)
    for _, char := range(s) {
        match, ok := matches[char]
        
        // If !ok then it must be an opening char
        if len(stack) == 0 || !ok {
            stack = append(stack, char)
        } else {
            top := stack[len(stack)-1]
            if match == top {
                // Pop off the stack
                stack = stack[:len(stack)-1]
            } else {
                // Closing char out of place
                return false
            }
        }
    }
    
    if len(stack) == 0 {
        return true
    } else {
        // If there is anything left on the stack
        // then it wasn't closed
        return false
    }
}
