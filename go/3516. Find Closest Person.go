func findClosest(x int, y int, z int) int {

    a:=math.Abs(float64(z-x))
    b:=math.Abs(float64(z-y))

    if a==b{
        return 0
    }else if a>b{
        return 2
    }else{
        return 1
    }
    
}