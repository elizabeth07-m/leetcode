func canBeTypedWords(text string, brokenLetters string) int {

    

    text1:=strings.Split(text," ")

    if brokenLetters==""{
        return len(text1)
    }
    brokenLetters1:=[]rune(brokenLetters)

   
 count:=0
    for _,v :=range text1{
       
        for j,val :=range brokenLetters1{
            if  strings.ContainsRune(v,val){
                break
            }else if len(brokenLetters)-1 == j{
              count++


            }
        }

    }
    return count
    
}