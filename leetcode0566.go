func matrixReshape(mat [][]int, r int, c int) [][]int {
    n, m := len(mat), len(mat[0])
    if r*c != n*m {
        return mat
    }
    res := make([][]int, r)
    for i := 0; i < m*n; i++ {
        if len(res[i/c]) == 0 {
            res[i/c] = make([]int, c)
        }
        res[i/c][i%c] = mat[i/m][i%m]
    }
    return res
}