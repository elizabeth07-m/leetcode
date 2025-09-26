func triangleNumber(nums []int) int {

    count:=0
    sort.Ints(nums)

    for i:=2 ; i< len(nums) ; i++{
        l,r :=0,i-1

        for  l<r{
            if nums[l]+nums[r]>nums[i]{
                count+=(r-l)
                r--
            }else{
                l++
            }
        }
    }
    return count
    
}