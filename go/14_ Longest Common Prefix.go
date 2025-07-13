func longestCommonPrefix(strs []string) string {

    s := strs[0]
    var ans string

    for i:=0 ; i<len(s) ; i++ {
     
     for j :=0 ;j<len(strs);j++{
            if len(strs[j])<=i || strs[j][i] != s[i]{
             return ans

            }
        
        
     }
             ans+=string(s[i])


    }
    return ans
    
}