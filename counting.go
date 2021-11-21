package counting

func Sort(nums []int) {
	// Find max number in array 'nums'
	maxNum := func(nums []int) (max int) {
		max = nums[0]
		for i := range nums {
			if nums[i] > max {
				max = nums[i]
			}
		}
		return max
	}(nums)

	// Create a slice with size of maxNum + 1 filled with 0's
	c := make([]int, maxNum+1, maxNum+1)

	// count the elements in 'nums'
	for i := range nums {
		c[nums[i]]++
	}

	// fill the array with correct sorted values
	pos := 0
	for i := range c {
		r := pos
		for ; r < pos+c[i]; r++ {
			nums[r] = i
		}
		pos = r
	}
}