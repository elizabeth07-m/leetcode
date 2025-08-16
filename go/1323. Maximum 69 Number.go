func maximum69Number (num int) int {

    str:= strconv.Itoa(num)

    var ans []int

    ans1:=10

    sol:=0
     j:=0

    for i:=0 ;i <len(str) ; i++{
  

   if str[i]=='9'|| j==1{

   ans=append(ans,int(str[i]-'0'))
   }else{
  ans=append(ans,9)
  j++

   }

   fmt.Println(ans,ans1)
sol=sol*ans1
sol=sol+ans[i]
fmt.Println(sol,ans[i])


    }

    return sol
    
}