func mySqrt(x int) int {

    if x<2{
        return x
    }

var sqrt int
    for i,j:=0,x ;i<=j;{
        mid :=(i+j)/2

        ans := mid * mid

        if ans > x{
            j=mid-1
        }else if ans < x{
            sqrt = mid
            i=mid+1

        }else{
            sqrt=mid
            return sqrt
        }


    }

return sqrt
    }

     




