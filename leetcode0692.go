package main

import "sort"

func topKFrequent(words []string, k int) []string {
	wordTimes := make(map[string]int)
	for _, s := range words {
		wordTimes[s]++
	}
	res := []string{}
	for w := range wordTimes {
		res = append(res, w)
	}
	sort.Slice(res, func(i, j int) bool {
		return wordTimes[res[i]] > wordTimes[res[j]] || wordTimes[res[i]] == wordTimes[res[j]] && res[i] < res[j]
	})
	return res[:k]
}
