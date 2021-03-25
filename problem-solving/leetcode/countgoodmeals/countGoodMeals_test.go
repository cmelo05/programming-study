package countgoodmeals

import "testing"

func TestCountPairs(t *testing.T) {
	deliciousness := []int{1, 3, 5, 7, 9}
	expectedCount := 4

	actualCount := countPairs(deliciousness)

	if expectedCount != actualCount {
		t.Errorf("Expected %d, got %d", expectedCount, actualCount)
	}
}

func TestCountPairs1(t *testing.T) {
	deliciousness := []int{1, 1, 1, 3, 3, 3, 7}
	expectedCount := 15

	actualCount := countPairs(deliciousness)

	if expectedCount != actualCount {
		t.Errorf("Expected %d, got %d", expectedCount, actualCount)
	}
}

func TestCountPairs2(t *testing.T) {
	deliciousness := []int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}
	expectedCount := 53

	actualCount := countPairs(deliciousness)

	if expectedCount != actualCount {
		t.Errorf("Expected %d, got %d", expectedCount, actualCount)
	}
}

// [2160,1936,3,29,27,5,2503,1593,2,0,16,0,3860,28908,6,2,15,49,6246,1946,23,105,7996,196,0,2,55,457,5,3,924,7268,16,48,4,0,12,116,2628,1468]
