func intersection(nums1 []int, nums2 []int) []int {

    var ans []int

    for i:=0 ; i<len(nums1) ; i++{

        for j:=0 ; j<len(nums2) ;j++{

            if nums1[i]==nums2[j]{

                ans=append(ans,nums1[i])
                break


            }
        }
    }

    sort.Ints(ans)
    var ans1 []int
   

    for  i:=0;i<len(ans)-1 ;i++{

        if ans[i] != ans[i+1] {
        
        ans1=append(ans1,ans[i])
        }

    }

    if len(ans)>=1{
       ans1=append(ans1,ans[len(ans)-1])
    }

    

    return ans1
    
}