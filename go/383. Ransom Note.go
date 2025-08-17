func canConstruct(ransomNote string, magazine string) bool {

    runes1 :=[]rune(ransomNote)
     runes2 :=[]rune(magazine)

    for i:=0 ;i<len(runes1);i++{
         
     for j:=0 ;j<len(runes2);j++{

        if runes1[i]==runes2[j]{
            runes2[j]='0'
            break

        }else if j ==len(runes2)-1{
            return false
        }
        
     }
     

    }
    return true

    
}