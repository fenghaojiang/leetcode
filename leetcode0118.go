func generate(numRows int) [][]int {
    res := make([][]int, 0)
    for i := 0; i < numRows; i++ {
        temp := make([]int, i+1)
        temp[0] = 1
        temp[i] = 1
        for j := 1; j < i; j++ {
            temp[j] = res[i-1][j] + res[i-1][j-1]
        }
        res = append(res, temp)
    }
    return res
}