func quickSort(arr []int, left, right int) {
	if left > right {
		return
	}
	i, j := left, right
	standard := arr[i]
	for i < j {
		for i < j && arr[j] > standard {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] < standard {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = standard
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}