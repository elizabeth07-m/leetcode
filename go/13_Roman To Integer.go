func romanToInt(s string) int {
    k := 0
    p := ""

    for i := len(s) - 1; i >= 0; i-- {
        ch := string(s[i]) 

        if ch == "I" && (p == "V" || p == "X") {
            k -= 1
        } else if ch == "X" && (p == "L" || p == "C") {
            k -= 10
        } else if ch == "C" && (p == "D" || p == "M") {
            k -= 100
        } else if ch == "I" {
            k += 1
        } else if ch == "V" {
            k += 5
        } else if ch == "X" {
            k += 10
        } else if ch == "L" {
            k += 50
        } else if ch == "C" {
            k += 100
        } else if ch == "D" {
            k += 500
        } else if ch == "M" {
            k += 1000
        }

        p = ch
    }

    return k
}
