func fizzBuzz(n int) []string {
	var answer []string
	var a string
	for value := 1; value <= n; value++ {
		if value%3 == 0 && value%5 == 0 {
			a = "FizzBuzz"

		} else if value%3 == 0 {
			a = "Fizz"

		} else if value%5 == 0 {
			a = "Buzz"

		} else {
			a = strconv.Itoa(value)

		}
		answer = append(answer, a)
	}
	return answer
}

func twoSum(nums []int, target int) []int {
var answer []int
        for i := 0 ; i <= len(nums)-1 ; i++{
         for j := i+1 ; j <= len(nums)-1 ; j++{
          if  nums[i]+nums[j] == target{
            answer = append (answer,i)
             answer = append (answer,j)

              return answer
          }

        }
    }
    return answer
}



func categorizeBox(length int, width int, height int, mass int) string {
    var bulky,heavy,answer string
    if (length >= 10000 || width >= 10000 || height >= 10000) ||  (length*width*height) >= 1000000000  {
        bulky="Bulky"
    }
    if mass >= 100{
        heavy="Heavy"
    }

    switch {
        case bulky == "Bulky" && heavy == "Heavy" : {
          answer="Both"

        }
                case bulky == "Bulky" && heavy == "Heavy" : {
          answer="Both"

        }
        case bulky != "Bulky" && heavy != "Heavy" : {
          answer="Neither"

        }
                case bulky == "Bulky" && heavy != "Heavy" : {
          answer="Bulky"

        }
                case bulky != "Bulky" && heavy == "Heavy" :{
          answer="Heavy"

        }



    }
    return answer
    
}