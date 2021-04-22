
func majorityElement(nums []int) []int {
	res := make([]int, 0)
	if nums == nil || len(nums)  == 0 {
		return res
	}
	candidate1 := nums[0]
	candidate2 := nums[0]
	var cnt1, cnt2 int 
	for _, v := range nums {
		if candidate1 == v {
			cnt1++
			continue
		}
		if candidate2 == v {
			cnt2++
			continue
		}
		if cnt1 == 0 {
			candidate1 = v
			cnt1++
			continue
		}
		if cnt2 == 0 {
			candidate2 = v
			cnt2++
			continue
		}
		cnt1--
		cnt2--
	}
	cnt1, cnt2 = 0, 0
	for _, v := range nums {
		if candidate1 == v {
			cnt1++
		} else if candidate2 == v {
			cnt2++
		}
	}
	l := len(nums) / 3 
	if cnt1 > l {
		res = append(res, candidate1)
	}
	if cnt2 > l {
		res = append(res, candidate2)
	}
	return res
}