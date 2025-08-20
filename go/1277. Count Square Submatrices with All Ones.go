func countSquares(matrix [][]int) int {
    
    for row := 1; row < len(matrix); row++ {
        for col := 1; col < len(matrix[0]); col++ {
            if matrix[row][col] == 0 {
                continue
            }
            matrix[row][col] = min(matrix[row-1][col-1], min(matrix[row-1][col], matrix[row][col-1])) + 1
        }
    }

    ans := 0
    for row := 0; row < len(matrix); row++ {
        for col := 0; col < len(matrix[0]); col++ {
            ans += matrix[row][col]
        }
    }
    return ans
}