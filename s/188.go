package s 

func maxProfit(k int, prices []int) int {
  n := len(prices)
  if n == 0 {
    return 0
  }
  if k >= n/2 {
    return fastpath(prices)
  }
  
  profits := make([][]int,k+1)
  for t := 0 ; t < k+1 ; t++ {
    profits[t] = make([]int, n)
  }
  
  for t := 1 ; t < k+1 ; t++ {
    curMax := -1 << 63
    for d := 1 ; d < n ; d++ {
      curMax = max(curMax, profits[t-1][d-1] - prices[d-1])
      profits[t][d] = max(profits[t][d-1], curMax + prices[d])
    }
  }
  return profits[len(profits)-1][len(profits[0])-1]
}

func fastpath(prices []int) int {
  p := 0 
  for i := 1 ; i < len(prices) ; i++ {
    p += max(0, prices[i]-prices[i-1])
  }
  return p
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}