package arrays

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//func sumAll(numbersToSum ...[]int) []int {
/*lenghtOfNumbers := len(numbersToSum)
sums := make([]int, lenghtOfNumbers)

for i := 0; i < lenghtOfNumbers; i++ {
	sums[i] = sum(numbersToSum[i])
}
return sums*/
//}

func sumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, sum(numbers))
	}

	return sums
}
