package palindromeNumber

import "strconv"

func isPalindrome(x int) bool {
	nums := strconv.Itoa(x)

	j := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[j] {
			return false
		}
		j--
	}

	return true
}
