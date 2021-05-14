package main

func intToRoman(num int) string {
	var charValues = []struct {
		Charater string
		Values   int
	}{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	var res []byte
	for _, v := range charValues {
		for num >= v.Values {
			res = append(res, []byte(v.Charater)...)
			num -= v.Values
		}
		if num == 0 {
			break
		}
	}
	return string(res)
}
