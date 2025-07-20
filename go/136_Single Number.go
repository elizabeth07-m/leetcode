// func singleNumber(nums []int) int {


//     for i:=0 ; i<len(nums) ; i++{
//         for j:=0 ;j <len(nums) ;j++{
//             if nums[i]==nums[j]&&i!=j{
//                 break

//             }else if j==len(nums)-1{
//                 return nums[i]
//             }

//         }
        
//     }
//     return 0
    
// }

func singleNumber(nums []int) int {

    sort.Ints(nums)
    for i:=0 ; i<len(nums)-1 ; i+=2{
        if nums[i]!=nums[i+1]{
            return nums[i]
        }
            
        
    }
    return nums[len(nums)-1]
    
}


















