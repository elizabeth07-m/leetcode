func intersect(nums1 []int, nums2 []int) []int {

      var ans []int

    for i:=0 ; i<len(nums1) ; i++{

        for j:=0 ; j<len(nums2) ;j++{

            if nums1[i]==nums2[j]{

                nums2[j]=-1
                ans=append(ans,nums1[i])
                break


            }
        }
    }

return ans
    
}