
func minimumArea(grid [][]int) int {
	leftMin := math.Inf(1)
	rightMax := 0
	mindepth := len(grid)
    maxdepth := 0
	for index, innerArr := range grid {
        // fmt.Println("Index is",index)
		for innerIndex, value := range innerArr {
			if value == 1 {
				if float64(innerIndex) <= leftMin {
					leftMin = float64(innerIndex)
				}
				if innerIndex >= rightMax {
					rightMax = innerIndex
				}
                if index >= maxdepth{
                    maxdepth = index
                }
                if index <= mindepth{
                    mindepth = index
                }
 
			} 
		}

	}

	width := rightMax - int(leftMin) + 1
    depth := maxdepth - int(mindepth) + 1
	return width * depth
}