func maxFreqSum(s string) int {
   
   var alp [26]int

   for i:=0;i<len(s);i++{
    alp[s[i]-'a']++
   }

   maxvow,maxcon:=0,0

    for i,v := range alp {
        if i==0||i==4||i==8||i==14||i==20{
            if maxvow<alp[i]{
                maxvow=v
            }

        }else if  maxcon <alp[i]{
                maxcon=v
            }
   }
   return maxvow+maxcon

}