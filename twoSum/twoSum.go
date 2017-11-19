package twoSum

func twoSum(nums []int, target int) []int {
    var result []int
    length := len(nums)
    for i := 0; i < length; i++ {
        complement := target - nums[i]
        for j := i + 1; j < length; j++ {
            if nums[j] == complement {
                result = []int{i, j}
                break
            }
        }
    }
    return result
}
