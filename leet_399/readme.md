[题目链接](https://leetcode-cn.com/problems/evaluate-division/description/)

### 题解
已知a/b和b/c，问a/c。根据数学公式很容易想到，a/c = a/b * b/c。也就是说，要求x/y，如果x/y的结果无法直接得出，那么如果找到x/z和z/y，那么它们的积就是x/y。很容易就能看出来这其实是道图论的题目。每个询问(x,y)实际上是在问x,y的连通性，如果不连通，就是-1.0，如果连通，则是路径上的值做累乘。

那么，已知部分点的连接关系，怎么求出所有点的连接关系呢？——Floyed算法。

Floyed求最短路简单讲就是，枚举每个节点k，再枚举两个节点i,j，如果`edge[i][k]+edge[k][j] < edge[i][j]`，则更新`edge[i][j]`。也就是说，枚举每个节点k，更新所有可以用k作为中转走到的点。代码非常简单，就是简单的三重循环：
```
for k in [x..y]
    for i in [x..y]
        for j in [x..y]
            if t:=dist[i][k]+dist[k][j]; t > dist[i][j]:
                dist[i][j] = t
```
本题只需要把比较逻辑做个小改动，只要`dist[i][k]`右边且`dist[k][j]`有边，则`dist[i][j] = dist[i][k]*dist[k][j]`
Floyed特别适用于稠密图中求所有点的最短距离，O(n^3)的复杂度，同时可以处理不带环的有负向边的图