func missingNumber(nums []int) int {

n:=len(nums)

total:=n*(n+1)/2
sum:=0

for _,val:=range nums{
    sum+=val

}

return total-sum

}
