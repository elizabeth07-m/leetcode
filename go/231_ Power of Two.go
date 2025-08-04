func isPowerOfTwo(n int) bool {
    if n < 1 {
        return false
    }

    a := 1
    for a <= n {
        if a == n {
            return true
        }
        a *= 2
    }
    return false
}
