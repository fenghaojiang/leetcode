func maxCount(m, n int, ops [][]int) int {
    mina, minb := m, n
    for _, op := range ops {
        mina = min(mina, op[0])
        minb = min(minb, op[1])
    }
    return mina * minb
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}