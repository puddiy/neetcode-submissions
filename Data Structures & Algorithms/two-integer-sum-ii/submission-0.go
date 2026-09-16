func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, num := range numbers {
		diff := target - num
		if j, ok := m[diff]; ok {
			return []int{j+1, i+1}
		}
		m[num] = i
	}

	return []int{}
}
