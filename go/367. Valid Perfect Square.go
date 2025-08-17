func isPerfectSquare(num int) bool {
    if num <=1 {
        return true 
    }

    i, j := 2, num/2
    for i <= j {
        mid := (i + j) / 2
        sq := mid * mid

        if sq == num {
            return true
        } else if sq < num {
            i = mid + 1
        } else {
            j = mid - 1
        }
    }
    return false
}
