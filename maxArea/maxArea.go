package maxArea

//func MaxArea(height []int) int {
//	max := 0
//	for i, v := range height {
//		for j, v2 := range height[i+1:] {
//			area := 0
//			if v2-v > 0 {
//				area = v * (j + 1)
//			} else {
//				area = v2 * (j + 1)
//			}
//			if area > max {
//				max = area
//			}
//		}
//	}
//	return max
//}

func MaxArea(height []int) int {
	max := 0
	length := len(height)
	for i := 0; i < length; i++ {
		start := height[i]
		area := 0
		for j := length - 1; j > 0; j-- {
			end := height[j]
			if end >= start {
				area = start * (j - i)
				if area > max {
					max = area
				}
				break
			} else {
				area = end * (j - i)
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}
