func searchInsert(nums []int, target int) int {
    var i int
    for i,_ =range nums{
        if target <= nums[i]{
            return i
        }
    }
    return i+1
    
}