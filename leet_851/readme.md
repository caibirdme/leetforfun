题目：

[原题链接](https://leetcode-cn.com/contest/weekly-contest-88/problems/loud-and-rich/)

在一组 N 个人（编号为 0, 1, 2, ..., N-1）中，每个人都有不同数目的钱，以及不同程度的安静（quietness）。
为了方便起见，我们将编号为 x 的人简称为 "person x "。
如果能够肯定 person x 比 person y 更有钱的话，我们会说 richer[i] = [x, y] 。注意 richer 可能只是有效观察的一个子集。
另外，如果 person x 的安静程度为 q ，我们会说 quiet[x] = q 。
现在，返回答案 answer ，其中 answer[x] = y 的前提是，在所有拥有的钱不少于 person x 的人中，person y 是最喧闹的人（也就是安静值 quiet[y] 最小的人）。

输入：richer = [[1,0],[2,1],[3,1],[3,7],[4,3],[5,3],[6,3]], quiet = [3,2,5,4,6,1,7,0]
输出：[5,5,2,5,4,5,6,7]

题解：

这道题模型也很明显，就是在图上进行简单的DP。
首先建图，对于[x,y]（x比y富有），就建立一条y到x的有向边。
所以最终建好的图中a到b有边就表示b比a有钱，更进一步地说，从a开始可以走到的点都比a有钱。

所以问题转化成了在一个有向图中，从a出发，找到quiet值最小的点。

用dist[k]表示从k出发能走到的quiet值最小的点，那么就有：

dist[k] = min(quiet[dist[i]]) connect[k][i] == true