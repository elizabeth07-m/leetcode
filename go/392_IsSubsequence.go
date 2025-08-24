func isSubsequence(s string, t string) bool {
    if len(s)==0{
        return true
    }

    for i,j:=0,0 ;i <len(s)&& j<len(t); {

        if s[i]==t[j] && i==len(s)-1{
           return true
        }else if s[i]==t[j]{
            i++
            j++
        }else {
            j++
        }
      
    }
    return false
    
}