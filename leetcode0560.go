func subarraySum(nums []int, k int) int {
    var count, pre int
    m := make(map[int]int)
    m[0] = 1
    for i := range nums {
        pre += nums[i]
        if _, ok := m[pre-k]; ok {
            count += m[pre-k]
        }
        m[pre] += 1
    }
    return count
}