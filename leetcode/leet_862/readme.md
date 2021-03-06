### 题目大意
[题目链接](https://leetcode-cn.com/contest/weekly-contest-91/problems/shortest-subarray-with-sum-at-least-k/)

一个int数组A，找一个连续的子数组，其和大于等于K，问：子数组长度最短是多少。

### 错误的解法
由于题目很短，我低估了他，导致想法比较简单，就是二分答案。利用滑动窗口找到有没有长度为L的子数组能够使得和大于等于K，如果有，则缩小查询范围。复杂度是NlogN。**但是这个思路有Bug！**，Bug点在于：
> 如果有多个解，能够二分答案去验证，必须保证如果M是一个解，那么t>M必然也是合法解。

仔细想一下，假如所有长度为M的子数组，其和都小于K，如果用二分法，我们这时就该考虑验证比M更长的子数组。但是，事实上比M更短的子数组，其和反而可能大于等于K，比如：`[1 -1 1 1]`
长度为3的子数组和为1，但是长度为2的子数组最大的和为2。
因此这里用二分法来验证答案是错误的！一定要注意答案的单调性。

### 正确的解法
这道题其实不复杂，首先我们求出所有的前项和sum[i]
* 如果`sum[j]-sum[i] >= k (j>i)`，那么最优解一定是最靠近j的那一个i，这样才能保证j-i最小。
* 如果`sum[j]-sum[i] < k`，只有在前面找个比sum[i]更小的sum[t]，sum[j]-sum[t]才有可能>=k。
先看下最朴素的N^2的做法：
```
for j:=sum.len-1; j>=1; j-- {
    for i:=j-1; i>=0; i-- {
        if sum[j]-sum[i]>=k {
            ans = min(ans, i-j)
            break
        }
    }
}
```
我们这里内层枚举了所有的sum[i]，但根据上面第二条，如果`sum[j]-sum[i]<k`，那么只需要向前枚举比sum[i]更小的值。如果有一个list，里面保存了一个单调的sum[i]，这样就可以显著的减少我们的枚举量。那么怎么来维护这个单调队列呢？根据上述的第一条，固定子数组右端点j不动，如果对单调队列dq有`sum[i]<=sum[dq.Last()]`，那么答案一定不可能是dq.Last()，因为如果sum[j]-sum[dq.Last()] >= k ，那么对于i来说，这个不等式也是成立的，且`j-i < j.dq.Last()`。因此这时可以把dq.Last()一直删除，直到sum[i]>sum[dq.Last()]。同时，如果`sum[j]-sum[dq.Front()]>=K`，这时可以记录下答案，然后删除掉dq.Front()，因为随着j的增大，j-dq.Front()的值只会更大。

当然，求前项和通常还需要注意，sum[0]表示没有项的和也就是0，而不是sum[0]=A[0]，应该是sum[1]=A[0]