func minNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        return fmt.Sprintf("%d%d", nums[i], nums[j]) < fmt.Sprintf("%d%d", nums[j], nums[i])
    })
    var res string 
    for i := range nums {
        res += fmt.Sprintf("%d", nums[i])
    }
    return res
}