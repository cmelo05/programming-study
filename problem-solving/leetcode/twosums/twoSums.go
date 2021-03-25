package twosums

func twoSum(nums []int, target int) []int {
	sum := []int{0, 0}

	for i, v1 := range nums {
		for j, v2 := range nums {
			if i != j && (v1+v2 == target) {
				sum[0] = j
				sum[1] = i
			}
		}
	}

	return sum
}
