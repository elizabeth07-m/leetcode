func findDisappearedNumbers(nums []int) []int {
    result := make([]int, len(nums))
    for _, v := range nums{
        result[v-1] = v
    }
    fmt.Println(result)
    insertPosition := 0

    for i := 0; i < len(nums); i++{
        if result[i] == 0{
            result[insertPosition] = i+1
            insertPosition++
        }
    }
    fmt.Println(result)
    return result[:insertPosition]
}