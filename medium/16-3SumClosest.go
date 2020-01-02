func threeSumClosest(nums []int, target int) int {
    closestDist := math.Inf(1)
    closestVal := 0
    sort.Ints(nums)
    
    for index := 0; index < len(nums)-2; index++ {
        left := index+1
        right := len(nums)-1
        
        for left < right {
            sum := nums[index] + nums[left] + nums[right]
            dist := math.Abs(float64(target - sum))
        
            if dist < closestDist {
                closestDist = dist
                closestVal = sum
            }

            if sum > target {
                right--
            } else {
                left++
            }
        }
    }
    
    return closestVal
}
