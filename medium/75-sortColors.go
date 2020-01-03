func sortColors(nums []int)  {
    if len(nums) <= 1 {
        return
    }
    
    pos := 0
    
    // These are the indexes to put 0s and 2s into
    left, right := 0, len(nums)-1
    
    for pos <= right {
        // Move it as far left as possible
        if nums[pos] == 0 && left != pos {
            temp := nums[left]
            nums[left] = nums[pos]
            nums[pos] = temp
            left++
        } else if nums[pos] == 2 && right != pos {
            // Move it as far right as possible
            temp := nums[right]
            nums[right] = nums[pos]
            nums[pos] = temp
            right--
        } else {
            // Nothing to swap, advance
            pos++
        }
    }    
}
