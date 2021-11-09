func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    var m, n = len(matrix), len(matrix[0])
    var i, j = m - 1, 0
    for i >= 0 && j < n {
        if matrix[i][j] < target {
            j++
        } else if matrix[i][j] > target {
            i--
        } else {
            return true
        }
    }
    return false
}
