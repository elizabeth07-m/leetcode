func removeElement(nums []int, val int) int {
var k int
    for i:=0 ;i<=len(nums)-1;i++{
        if val !=nums[i]{
            nums[k]=nums[i]
            k++
        }
    }
   
    return k
    
}