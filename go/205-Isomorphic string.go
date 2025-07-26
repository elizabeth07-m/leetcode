func isIsomorphic(s string, t string) bool {
    if len(s)!=len(t){
        return  false
    }

    for i:=0 ;i<len(s) ;i++{

        for j:=i+1;j<len(s);j++{

            if s[i]==s[j]{

                if t[i]!=t[j]{
                    return false
                }


            }else{
                  if t[i]==t[j]{
                    return false
                }

            }


        }



    }

    return true

    
}