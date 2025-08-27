
func minimumArea(grid [][]int) int {

    left:=0
    right:=0
    top:=0
    bott:=0
    first:=0

    for i:=0 ;i<len(grid);i++{
        for j:=0 ;j<len(grid[0]);j++{

          if grid[i][j]==1 && first==0{
              left = j
              right=j
              top =i
              bott =i
              first++
          } else if grid[i][j]==1{
              left = min(left,j)
              right=max(right,j)
              top =min(top,i)
              bott =max(bott,i)
          } 


        }
        
    }
    fmt.Println(left,right,top,bott)
    return ((right-left)+1)*((bott-top)+1)












	// leftMin := math.Inf(1)
	// rightMax := 0
	// mindepth := len(grid)
    // maxdepth := 0
	// for index, innerArr := range grid {
    //     // fmt.Println("Index is",index)
	// 	for innerIndex, value := range innerArr {
	// 		if value == 1 {
	// 			if float64(innerIndex) <= leftMin {
	// 				leftMin = float64(innerIndex)
	// 			}
	// 			if innerIndex >= rightMax {
	// 				rightMax = innerIndex
	// 			}
    //             if index >= maxdepth{
    //                 maxdepth = index
    //             }
    //             if index <= mindepth{
    //                 mindepth = index
    //             }
 
	// 		} 
	// 	}

	// }

	// width := rightMax - int(leftMin) + 1
    // depth := maxdepth - int(mindepth) + 1
	// return width * depth
}