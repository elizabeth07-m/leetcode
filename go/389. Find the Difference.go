func findTheDifference(s string, t string) byte {
   
var ans1 int
var  ans2 int 

  for i,j :=0,0 ;i<len(s)||j<len(t) ; i,j=i+1,j+1 {
    
    if i<len(s){
    ans1+=int(s[i])
    }
    if j<len(t){
     ans2+=int(t[j])
    }
  }

   return byte(ans2-ans1)

    
}
// func findTheDifference(s string, t string) byte {
//     runes1 :=[]rune(t)
//     runes2 :=[]rune(s)

//   for i:=0 ;i<len(runes1) ;i++{

//     for j:=0 ; j<len(runes2);j++{

//         if runes1[i]==runes2[j]{
//            runes2[j]=0
//             break
//         }else if j==len(runes2)-1{
//             fmt.Println(runes1[i])
//             return t[i]
//         }
//     }
//   }
//    return t[0]

    
// }