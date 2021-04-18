func generateParenthesis(n int) []string {
    res := make([]string, 0)
    var dfs func(lRemain, rRemain int, s string)
    dfs = func(lRemain, rRemain int, s string) {
        if 2*n == len(s) {
            res = append(res, s)
            return
        }
        if lRemain > 0 {
            dfs(lRemain-1, rRemain, s+"(")
        }
        if lRemain < rRemain {
            dfs(lRemain, rRemain-1, s+")")
        }
    }
    dfs(n, n, "")
    return res
}