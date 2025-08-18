func firstUniqChar(s string) int {

    ans := make(map[string]int)

    for i:=0 ; i<len(s) ; i++{

     if _,ok :=ans[string(s[i])];ok{
       ans[string(s[i])]=-1
     }else{
        ans[string(s[i])]=i
     }

    }
    fmt.Println(ans)


     for i:=0 ; i<len(s) ; i++{

       for k,v := range ans{
        if k==string(s[i])&&v!=-1{
            return v
        }
        
    }


     }

  
    return -1
    
}