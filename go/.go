func maxFrequencyElements(nums []int) int {
    maxFrequencyByElement := make(map[int]int)
    totalMaxFrequency := 0
    elementsWithTotalMaxFrequency := 0

    for _, e := range nums {
        maxFrequencyByElement[e]++
    }

    for _, v := range maxFrequencyByElement {
        if v == totalMaxFrequency {
            elementsWithTotalMaxFrequency++
        } else if v > totalMaxFrequency {
            totalMaxFrequency = v
            elementsWithTotalMaxFrequency = 1
        }
    }

    return totalMaxFrequency * elementsWithTotalMaxFrequency
}