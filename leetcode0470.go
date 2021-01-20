func rand10() int {
	for {
		ans := (rand7()-1)*7 + rand7()
		if ans <= 40 {
			return 1 + ans%10
		} //生成范围1-49，取1-40
		ans = (ans-40-1)*7 + rand7() //大于40的9个数再次生成随机数
		if ans <= 60 {
			return 1 + ans%10
		} //生成范围1-63，取1-60
		ans = (ans-60-1)*7 + rand7() //大于60的3个数再次生成随机数
		if ans <= 20 {
			return 1 + ans%10
		} //生成范围1-21，取1-20
	}
}