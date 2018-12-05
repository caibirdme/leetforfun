### 题解

[题目链接](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)

股票交易，给定每日股价，必须卖了再买，最多交易K次，求最大利润

### 暴力（N^3logN，TLE）
用`f[i][j]`表示“前i天进行j次交易且第j次交易是在第i天完成的”的最大利润。也就是说枚举最后一次交易，`f[i][j] = Max{f[k][j-1] + profit(k+1,i)}`。k是上次交易的截止日期，那么本次交易的收益就是profit(k+1,i)，且`profit(k+1,i)= price[i]- min{price[k+1..i-1]}`

求`min{price[k+1..i-1]}`可以用线段树

结论：太复杂且严重超时

### DP
`f[i][j][k=0|1]`表示前i天一共进行不超过j次交易且最后一天不卖(k=0)|卖(k=1)的最大收益。我们分情况讨论：
* 最后一天不卖：`f[i][j][0] = max(f[i-1][j][0], f[i-1][j][1])`，这个很好理解
* 最后一天卖：`f[i][j][1] = max(f[i-1][j][1], f[i-1][j-1][0]) - price[i-1] + price[i]`

由于转移只涉及到j-1，因此可以使用01滚动或者倒序枚举节约一维空间