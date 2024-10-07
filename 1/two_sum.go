package twoSum

func twoSum(nums []int, target int) []int {
	for i, first := range nums {
		for j, second := range nums[:i] {
			if first+second == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
