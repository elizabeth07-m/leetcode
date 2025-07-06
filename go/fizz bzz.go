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