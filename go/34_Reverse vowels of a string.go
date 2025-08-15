func reverseVowels(s string) string {
   
   runes :=[]rune(s)

   vowels :="AEIOUaeiou"

   var pos []int

   for i,v := range runes{
    
    if strings.ContainsRune(vowels,v){
        pos=append(pos,i)
    }


   }

  for i,j:=0,len(pos)-1 ; i<j ;i,j=i+1,j-1{

    runes[pos[i]], runes[pos[j]]= runes[pos[j]], runes[pos[i]]


  }

  return string(runes)

   


}
