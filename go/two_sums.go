func twoSum(nums []int, target int) []int {
var answer []int
        for i := 0 ; i <= len(nums)-1 ; i++{
         for j := i+1 ; j <= len(nums)-1 ; j++{
          if  nums[i]+nums[j] == target{
            answer = append (answer,i)
             answer = append (answer,j)

              return answer
          }

        }
    }
    return answer
}