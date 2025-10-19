import (
    "strings"
)

func findLexSmallestString(s string, a int, b int) string {
    visited := map[string]bool{}
    minStr := s
    dfs(s, a, b, visited, &minStr)
    return minStr
}

func dfs(curr string, a int, b int, visited map[string]bool, minStr *string) {
    if visited[curr] {
        return
    }
    
    visited[curr] = true
    *minStr = min(*minStr, curr)

    dfs(add(curr, a), a, b, visited, minStr)
    dfs(rotate(curr, b), a, b, visited, minStr)
}

func add(s string, a int) string {
    res := []string{}
    for i := range s {
        if i%2 == 0 {
            res = append(res, string(s[i]))
        } else {
            val, _ := strconv.Atoi(string(s[i]))
            val = (val + a)%10
            res = append(res, strconv.Itoa(val))
        }
    }
    return strings.Join(res, "")
}

func rotate(s string, b int) string {
    b = b % len(s)
    bs := []byte(s)
    reverse(bs)
    reverse(bs[:b])
    reverse(bs[b:])
    return string(bs)
}

func reverse(s []byte) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}