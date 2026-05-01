func productExceptSelf(nums []int) []int {
    // naive
    output := make([]int, len(nums))

    p := nums[0]
    zeroes := map[int]int{}
    zCount := 0

    for i := 1; i < len(nums); i++ {
        if nums[i] == 0 {
            zeroes[i] = 0
            zCount++
            continue
        }

        p *= nums[i]    
    }

    for i := 0; i < len(nums); i++ {
        if zCount >= 2 {
            output[i] = 0
        }

        if nums[i] != 0 && zCount > 0 {
            output[i] = 0
        }

        if nums[i] != 0 && zCount == 0 {
            output[i] = p / nums[i]
        }

        if nums[i] == 0 && zCount == 1 {
            output[i] = p
        }
    }

    return output
}
