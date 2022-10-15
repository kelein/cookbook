package tests

func sum(number []int) int {
	count := 0
	for _, n := range number {
		count += n
	}
	return count
}

func sumAll(nums ...[]int) []int {
	sumarr := make([]int, len(nums))
	for i, num := range nums {
		sumarr[i] = sum(num)
	}
	return sumarr
}

func sumAllTails(nums ...[]int) []int {
	sums := make([]int, len(nums))
	for i, num := range nums {
		if len(num) > 0 {
			sums[i] = sum(num[1:])
		} else {
			sums[i] = 0
		}
	}
	return sums
}
