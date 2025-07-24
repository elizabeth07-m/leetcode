func isHappy(n int) bool {

    a:= strconv.Itoa(n)
    sum:=0
    seen := make(map[int]bool)

    for i:=0 ;i< len(a) ;i++{

        g,_:=strconv.Atoi(string(a[i]))
         ans:=g*g

     sum+=ans
     if sum == 1 && i ==len(a)-1{
        return true
     }else if i ==len(a)-1{
        if seen[sum]{

            return false
        }else{
            seen[sum]=true
        }
        a= strconv.Itoa(sum)
        sum=0
        i=-1
        
     }


    }
    return false
    
}