package main

func removeDuplicates(nums []int) int {
	unique := make(map[int]int)
	count := 0
	for _, i := range nums {
		if _, ok := unique[i]; !ok {
			unique[i] = 1
			nums[count] = i
			count += 1
		}
	}
	nums = nums[:count]
	return len(nums)
}
