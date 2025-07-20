func mySqrt(x int) int {


    for i:=0 ;i<=x ; i++{
        a:=i*i
        if a==x{
            return i
        }else if a > x{
            return i-1
        }
    }
    return 1
}