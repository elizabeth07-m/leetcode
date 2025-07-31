func containsNearbyDuplicate(nums []int, k int) bool {

    indexmap :=  make(map[int]int)

    for i :=0; i<len(nums) ; i++{

   
     if val,found:=indexmap[nums[i]];found && i-val<=k{
        return true
    }
      indexmap[nums[i]]=i
    }
    return false
   

}
