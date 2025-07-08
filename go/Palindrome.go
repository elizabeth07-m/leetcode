func isPalindrome(x int) bool {
    xstring := strconv.Itoa(x)
 
     arrayOfString := make([]string,len(xstring))

    for index,value := range xstring{
        arrayOfString[index]= string(value)
    }
    var answer []string

    for i := len(arrayOfString)-1 ;i>=0 ; i-- {
        answer =append(answer,arrayOfString[i])
    }
    ans := strings.Join(answer,"")
    

    if ans == xstring{
        return true
    }
    return false
    
}