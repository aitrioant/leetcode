package twoSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := [...]int{2, 7, 11, 15}
	target := 9

	indexes := twoSum(nums[:], target)

	result := 0
	for _, index := range indexes {
		result += nums[index]
	}
	assert.Equal(t, target, result)

}

func TestTwoSum2(t *testing.T) {
	nums := [...]int{3, 2, 4}
	target := 6

	indexes := twoSum(nums[:], target)

	result := 0
	for _, index := range indexes {
		result += nums[index]
	}
	assert.Equal(t, target, result)

}

func TestTwoSum3(t *testing.T) {
	nums := [...]int{3, 3}
	target := 6

	indexes := twoSum(nums[:], target)

	result := 0
	for _, index := range indexes {
		result += nums[index]
	}
	assert.Equal(t, target, result)

}
