func maxProfit(prices []int) int {
    min := math.MaxInt32
    var profit int    
    for i := range prices {
        if prices[i] < min {
            min = prices[i]
        }
        if profit < prices[i] - min {
            profit = prices[i] - min
        }
    }
    return profit
}