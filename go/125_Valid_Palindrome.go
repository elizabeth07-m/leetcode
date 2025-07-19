func isPalindrome(s string) bool {
    
    var a []byte
    for _,v:= range s {
        if (int(v)>=97 && int(v)<=122){
           a=append(a,byte(v))
        }else if(int(v)>=65 && int(v) <=90){
            b:= byte(int(v) + 32)
            a=append(a,b)
        }else if (int(v)>=48 && int(v)<=57){
             a=append(a,byte(v))

        }


    }
     
     i:=0
     j:=len(a)-1

     for i<j{
        if(a[i] != a[j]){
            return false
        }
        i++
        j--
     }
return true
}