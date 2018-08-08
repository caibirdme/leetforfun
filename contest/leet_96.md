# Contest 96

### 三维形体投影面积
[题目链接](https://leetcode-cn.com/contest/weekly-contest-96/problems/projection-area-of-3d-shapes/)
#### 题目描述
每个物体长宽高都是1，告诉三维坐标系里每个坐标点i (x,y)上有Ti个物体堆叠，求所有物体在每个面的投影面积。（原题是求总投影面积，实际上求出每个面的投影面积，三个加起来就行了）

#### 题解
* 最简单的，在xy平面里的投影面积，其实就是有多少个坐标上有物体
* 在xz平面上的投影：看看两个点 (x1,y1,z1), (x1,y2,z1)，这两个点在xz平面上的投影面积是什么？ 因为在xz平面上，所以从y轴来看，对于同样的x，不论y怎么变，只能看到最高的那个。也就是说，当x=x1时，求出所有(x1, yi)的最大值k1，投影就是k1。枚举所有x，求出所有kx即可。
* yz轴同理

很简单的题目，详见[AC代码](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_887/source.go)

### 救生艇
[题目链接](https://leetcode-cn.com/contest/weekly-contest-96/problems/boats-to-save-people/)

#### 题目描述
有N个人，每个人体重是people[i]。已知救生艇最大载重为limit，且每艘救生艇不论如何最多载两人。问需要多少救生艇。

#### 题解
这道题是个贪心，先把体重按从小到大排序。两个指针left,right，分别指向剩下的人中最轻和最重的。每个人都要上一艘救生艇，因此people[left]先选一艘。left选了一艘之后，这艘救生艇还可以载一个人，那肯定在能力范围内载一个越重的人越好，因此就是right了。如果people[left]+people[right]>limit，那right那位兄台只能自己做一艘船了（因为当前最轻的人和他一起坐都超重了）。

详见[AC代码](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_885/source.go)

### 索引处的解码字符串
[原题链接](https://leetcode-cn.com/contest/weekly-contest-96/problems/decoded-string-at-index/)

#### 题目描述
对于一个字符串`a2b2c2d2`和一个空串S，每当遇到字符c，就往S后面追加一个c。如果遇到一个数字t，就把S repeat t次。
对于`a2b2c2d2`，S最后就成了：aabaabcaabaabcdaabaabcaabaabcd。
给定这个字符串，问最后S第k个字符是什么

#### 题解
暴力的做法就是模拟S串的生成过程，然后输出第K位。但是这个不论是时间复杂度还是空间复杂度都太高。我们一起来想一想，纯模拟的方法到底哪些步骤浪费了时间或空间。首先我们可以把整个过程划分成两个状态：
* append
* repeat

append状态是说，上一次之前串repeat完成了，现在读到了字符，直接append到S后面。repeat状态是说，读到了一个数字，现在把串S repeat N次。串S要么处在append状态要么处在repeat状态。

那么如果已知串S，repeat N次，这时候我们真的需要开辟一个N*len(S)的空间来存新的串吗？是不是只需要存成(S,N)这样的tuple就好了？由于串是以len(S)的长度循环的，因此求第K位，等价于求第K%len(S)位。

想清楚了这个，我们就接着继续思考。要求整个串的第K位，实际上当S扩展到长度K时，它要么是通过append扩展的，要么是通过repeat扩展的。那我们怎么知道它是从append还是repeat扩展出来的呢？——记录最后一次repeat。怎么理解`记录最后一次repeat`？假设我们有一个sum[i]表示第i次repeat之后S串的总长度，接下来假设读到了连续T个字符，第T+1是数字M。我们很容易发现：
* 下一次循环节的长度是sum[i]+T
* 下一次repeat之后的总长度是(sum[i]+T)*M

假设`K>sum[i] && K<=sum[i]+T`，说明第K位是append状态扩展而来，这时直接输出新读到的第K-sum[i]个字符就行了。

假设`K>sum[i]+T && K<(sum[i]+T)*M`，说明第K位是repeat扩展得到的，循环节长度是sum[i]+T。由前面的推论得知，此时第K位是第K%(sum[i]+T)位是一样的，因此问题转化成了求S的第K%(sum[i]+T)位。

这就变成了一个相同子问题，我们通过递归就能求解了。再想一想递归的边界是什么？一个字符首先得经过append，才会进入后续的repeat状态，因此对于某个字符c，它第一次一定是由append扩展得到的。所以我们的递归只要有判断append的分支，是一定能到达递归边界的。

详见[AC代码](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_884/source.go)

### 细分图中的可到达结点
[原题链接](https://leetcode-cn.com/contest/weekly-contest-96/problems/reachable-nodes-in-subdivided-graph/)

#### 题解
这道题我的解法比较绕，官方的题解简单易懂。下面说说我的做法吧：
首先用最短路算法来先求出点0到其它所有点的最短路，我这里使用的是SPFA。求最短路有什么用呢？考虑以下case，假设从0出发，经过各种绕走到了3号点，耗费了8步。假设最多走10步，这就意味着从3号点继续走还剩两步可以走。但是如果从0走到3的最短路是5，那么意味着沿着最短路径走到3，此时从3出发还能再走5步，这意味着可达的点更多。
也就是说，如果某条边属于最短路径上的边，那么这条边上所有点一定是可达的，累加到答案ans。如果某条边不是最短路径上的边，那么这条边能走多长是多长。但是走完之后得记录这条边有多长已经是被走过了。

具体实现还有很多细节处理，详见[AC代码](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_886/source.go)

### 总结
这套题还是有点儿难度的，能在1个半小时内全部AC的都是大佬，我最后一道题加上调试花了差不多一个晚上…