func plusOne(digits []int) []int {

     
             for i:=len(digits)-1 ;i >=0 ;i--{

                if digits[i]!=9{
                    digits[i]=digits[i]+1
                    return digits
                }else{
                    digits[i]=0
                    if i==0{
                        digits=append([]int{1},digits...)
                        return digits
                    }
                } 
             }

      





    
        

   return digits
     
}