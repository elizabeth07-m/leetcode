func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

ans :=make(map[byte]int,len(s))

    for i:=0 ; i<len(s) ;i++{
        if _,ok :=ans[s[i]];ok{
            ans[s[i]]++
        }else{

        ans[s[i]]=1
        }
          if _,ok :=ans[t[i]];ok{
            ans[t[i]]--
        }else{

        ans[t[i]]=-1
        }


    }

      for _,val:=range ans{
        if val!=0{
            return false
        }
        }
      
 return true   
}
