Problem:

There are N network nodes, labelled 1 to N.

Given times, a list of travel times as directed edges times[i] = (u, v, w), where u is the source node, v is the target node, and w is the time it takes for a signal to travel from source to target.

Now, we send a signal from a certain node K. How long will it take for all nodes to receive the signal? If it is impossible, return -1.

Note:
N will be in the range [1, 100].
K will be in the range [1, N].
The length of times will be in the range [1, 6000].
All edges times[i] = (u, v, w) will have 1 <= u, v <= N and 1 <= w <= 100.


其实就是最短路的模板题。从K开始求出到所有点的最短路，如果某个点不可达则输出-1，否则输出最长的那条路径。
模板题可选算法挺多：
* floyd
* dijkstra
* spfa
* bellmanford

bellmanford可以算负权边的图，但是实际上可以通过平移把边转成整数，因此基本上只用spfa就够了。数据规模较大的稠密图可以选用dijkstra，选取点的时候用heap来选取一个最大值。这个实现稍微复杂一点。

spfa大部分情况下是最佳的选择