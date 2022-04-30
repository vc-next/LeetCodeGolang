package sum

//func TwoSum(nums []int, target int) []int {
//	for i, v := range nums {
//		for i2, v2 := range nums[i+1:] {
//			if v2 == target-v {
//				return []int{i, i2 + i + 1}
//			}
//		}
//	}
//	return []int{}
//}

func TwoSum(nums []int, target int) []int {
	table := map[int]int{}
	for i, v := range nums {
		minus, ok := table[target-v]
		if ok {
			return []int{minus, i}
		}
		table[v] = i
	}
	return []int{}
}
