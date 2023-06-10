package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {

	// init sumSlice
	sumSlice := []int{}

	for i := 0; i < len(numbersToSum); i++ {
		append(sumSlice, Sum(numbersToSum[i]))
		// & Trying to append value :)
	}

	// Need to Loop through each parameter slice.

	// Sum each slice, push that to sumSlice

	// Return sumSlice

	return nil
	// return []int{}
}
