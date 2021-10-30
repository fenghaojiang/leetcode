func kthDistinct(arr []string, k int) string {
    m := make(map[string]int)
    for i := range arr {
        m[arr[i]]++
    }
    var cnt int
    for i := range arr {
        if m[arr[i]] == 1 {
            cnt++
            if cnt == k {
                return arr[i]
            }
        }
    }
    return ""
}