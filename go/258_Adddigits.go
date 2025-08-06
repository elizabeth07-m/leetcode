func addDigits(num int) int {
ans:=0

if num<10{
    return num
}
   
    for num>9{

      ans=rec(num)
      num=ans



    }
    return ans
    
}

func rec (num int)int{

      ans:=0 
    quo :=0
    rem :=0
      for num>0{

        quo = num/10
        rem  =num % 10
        num=quo
        ans+=rem



    }
    return ans

}