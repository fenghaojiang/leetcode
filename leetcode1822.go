func arraySign(nums []int) int {
    var cnt int 
    for _, v := range nums {
        if v == 0 {
            return 0
        } 
        if v < 0 {
            cnt++
        }
    }
    if cnt % 2 == 0 {
        return 1
    } else {
        return -1
    }
}