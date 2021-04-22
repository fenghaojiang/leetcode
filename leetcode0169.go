package main

func majorityElement(nums []int) int {
	var cnt, major int
	for _, v := range nums {
		if cnt == 0 {
			major = v
		}
		if major == v {
			cnt++
		} else {
			cnt--
		}
	}
	return major
}
