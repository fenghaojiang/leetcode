func canConstruct(ransomNote string, magazine string) bool {
    ranNum := [26]int{}
    for _, r := range magazine {
        ranNum[r-'a']++
    }
    for _, r := range ransomNote {
        ranNum[r-'a']--
        if ranNum[r-'a'] < 0 {
            return false
        }
    }
    return true
}