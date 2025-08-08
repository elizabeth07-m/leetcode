func missingNumber(nums []int) int {

    sort.Ints(nums)

  

    for i:=0 ; i<len(nums);i++{
        if nums[i]!=i{
            return i
        }
    }
    return nums[len(nums)-1]+1
    
}