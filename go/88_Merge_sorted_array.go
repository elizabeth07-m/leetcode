func merge(nums1 []int, m int, nums2 []int, n int)  {


    for i,j := 0, m; j < n; i, j = i+1, j+1 {
    nums1[i]=nums2[j] 
 
    }
   
sort.Ints(nums1)
 
    
}