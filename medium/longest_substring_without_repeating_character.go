package main

func lengthOfLongestSubstring(s string) int {
	max1 := 0
	max2 := 0
	for i := 0; i < len(s); i++ {
		unique := make(map[uint8]int)
		for j := i; j < len(s); j++ {
			if _, ok := unique[s[j]]; !ok {
				unique[s[j]] = 1
				max1 += 1
			} else {
				break
			}
		}
		if max2 < max1 {
			max2 = max1
		}
		max1 = 0
	}
	return max2
}
