func majorityElement(nums []int) int {
    sort.Ints(nums)
 count:=1
    for i:=0 ;i < len(nums)-1 ;i++{
    

            if nums[i]== nums[i+1] {
                count++
                


            }
            if count>len(nums)/2{
                return nums[i]
            }else if nums[i]!= nums[i+1] {

                count=1
            }
        
      
    }
    return nums[0]
    
}