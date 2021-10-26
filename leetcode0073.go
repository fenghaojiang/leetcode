func setZeroes(matrix [][]int)  {
    n, m := len(matrix), len(matrix[0])
    var row0, col0 = false, false 
    for i := range matrix[0] {
        if matrix[0][i] == 0 {
            row0 = true
            break
        }
    }
    for i := range matrix {
        if matrix[i][0] == 0 {
            col0 = true
            break
        }
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    if row0 {
        for i := 0; i <  m; i++ {
            matrix[0][i] = 0
        }
    }
    if col0 {
        for i := 0; i <  n; i++ {
            matrix[i][0] = 0
        }
    }
}