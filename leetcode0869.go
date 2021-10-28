func countDigit(n int) [10]int {
    cnt := [10]int{}
    for n != 0 {
        cnt[n%10]++
        n /= 10
    }
    return cnt
}

var mapTwo map[[10]int]bool

func init() {
    mapTwo = make(map[[10]int]bool)
    for i := 1; i <= 1e9; i <<= 1 {
        mapTwo[countDigit(i)] = true
    }
}

func reorderedPowerOf2(n int) bool {
    return mapTwo[countDigit(n)]
}  


