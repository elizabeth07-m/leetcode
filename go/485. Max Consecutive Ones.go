func findMaxConsecutiveOnes(nums []int) int {

     a,b :=0,0

    for i,_:= range nums{

        if nums[i]==1{
            a++

        }else{
            if b<a{
            b=a
            }
            a=0
        }
          if b<a{
            b=a
            }
   
    }
    return b
    
}