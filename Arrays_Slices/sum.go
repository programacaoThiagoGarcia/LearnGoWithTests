package main

func sum(numbers []int) int {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, value := range numbersToSum {
		sums = append(sums, sum(value))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, value := range numbersToSum {
		tail := value[1:]
		sums = append(sums, sum(tail))
	}

	return sums
}
