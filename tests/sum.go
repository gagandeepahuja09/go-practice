package sum

func ints(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + ints(nums[1:]...)
}

// nums = []int{1, 2, 3, 4}
// sum(nums...) will convert it into multiple arguments
// nums ...{type} will convert it into a slice of type

// variadic functions
func Ints(nums ...int) int {
	return ints(nums...)
}
