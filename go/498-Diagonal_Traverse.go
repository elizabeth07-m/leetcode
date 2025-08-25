func findDiagonalOrder(mat [][]int) []int {
    res := make([]int, len(mat) * len(mat[0]))
    var key bool
    var i, j, in, jn = 0, 0, len(mat), len(mat[0])
    res[0] = mat[0][0]
    for t := 1; t < len(res); t++ {
        if key == false {
            if i-1 < 0 || j+1 >= jn{
                key = true
                if j+1 < jn {
                    j++
                } else {
                    i++
                }
            } else {
                i--
                j++
            }
        } else {
            if i+1 >= in || j-1 < 0{
                key = false
                if i+1 < in {
                    i++
                } else {
                    j++
                }
            } else {
                i++
                j--
            }
        }
        res[t] = mat[i][j]
    }
    return res
}