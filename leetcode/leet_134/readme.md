## 题解

### 怎么处理环？
一个最简单的方法是，把数组扩大一倍，也就是说假设数组是`[2,5,1]`，扩大一倍变成`[2,5,1,2,5,1]`。当然这只能用于处理一次环，如果多次就要想别的方法了…

### brute force
暴力方法就是枚举每个点k，把它当起点，然后一直走看能不能走回到点k。这种方法的复杂度是n^2。
有个小优化，如果一个点i，`gas[i]<cost[i]`，这个点肯定就能被排除掉。

### 怎么优化
根据上面那个优化case，当且仅当`gas[i]>=cost[i]`时，点i才可能是起点。因此我们可以用一个队列来保存这些**可能是起点**的点。
我们枚举一个起点i，然后开始往后走。在这个过程中，每当发现一个**可能是起点**的点，我们就将它加入队列。当走到点k发现无法继续了，说明以i为起点不行，这个时候由于我们按顺序记下了可能作为起点的点，因此直接从队列里取出点，看是否可行。
我们先计算出sumGas和sumCost，假设从点i走到点p还剩rest的油，如果把起点从i换成k，我们可以利用sumGas和sumCost求出: `rest = rest-(sumGas[k-1]-sumGas[i-1])+sumCost[k-1]-sumCost[i-1]`

因此这种方法的时间复杂度是`O(2n)`，提交记录显示击败100%。