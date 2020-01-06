func simplifyPath(path string) string {
    stack := make([]string, 0)
    parts := strings.Split(path, "/")
    
    for _, part := range(parts) {
        if part == "" || part == "." {
            continue
        }
        
        if part == ".." {
            // Pop
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            }
        } else {
            stack = append(stack, part)
        }
    }
    
    if len(stack) == 0 {
        return "/"
    } else {    
        var builder strings.Builder
        res := strings.Join(stack, "/")

        builder.WriteRune('/')
        builder.WriteString(res)

        // Remove the trailing slash
        return builder.String()
    }
}
