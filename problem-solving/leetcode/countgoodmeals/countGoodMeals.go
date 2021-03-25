package countgoodmeals

import "fmt"

func countPairs(deliciousness []int) int {
	count := 0

	frequencyList := make(map[int]int)

	for _, v := range deliciousness {
		value, ok := frequencyList[v]
		if ok {
			frequencyList[v] = value + 1
		} else {
			frequencyList[v] = 1
		}
	}

	for k1, v1 := range frequencyList {
		for k2, v2 := range frequencyList {
			sum := k1 + k2

			if k1 != k2 && sum > 0 && (sum == 1 || (sum&(sum-1)) == 0) {
				fmt.Printf("[%d, %d] \n", k1, k2)
				fmt.Printf("Sum: %d \n", sum)
				fmt.Printf("Frequency 1: %d \n", v1)
				fmt.Printf("Frequency 2: %d \n", v2)
				count = count + v1 + v2
			}
		}

	}

	return count
}

func countPairsUnefficient(deliciousness []int) int {
	count := 0

	for i, v1 := range deliciousness {
		for j := i + 1; j < len(deliciousness); j++ {
			sum := v1 + deliciousness[j]
			if i != j && sum > 0 && (sum == 1 || (sum&(sum-1)) == 0) {
				count++
				fmt.Printf("[%d, %d] \n", v1, deliciousness[j])
				fmt.Printf("Sum: %d \n", sum)
			}
		}
	}

	return count
}
