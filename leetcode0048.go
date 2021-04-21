func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l/2; i++ {
		for j := 0; j < (l+1)/2; j++ {
			matrix[i][j], matrix[l-j-1][i], matrix[l-i-1][l-j-1], matrix[j][l-i-1] = matrix[l-j-1][i], matrix[l-i-1][l-j-1], matrix[j][l-i-1], matrix[i][j]
		}
	}
}