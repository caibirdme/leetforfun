题意：

[原题链接](https://leetcode-cn.com/contest/weekly-contest-88/problems/maximize-distance-to-closest-person/)

在一排座位（ seats）中，1 代表有人坐在座位上，0 代表座位上是空的。
至少有一个空座位，且至少有一人坐在座位上。
亚历克斯希望坐在一个能够使他与离他最近的人之间的距离达到最大化的座位上。
返回他到离他最近的人的最大距离。

输入：[1,0,0,0,1,0,1]
输出：2

输入：[1,0,0,0]
输出：3

题解：

其实就是扫一遍，找到最长的连续的0有多长。如果最长的连续0在开头或者结尾，那么答案就是它，如果在中间，答案就是:
1. maxConsecutiveZero/2 (maxConsecutiveZero为偶数)
2. (maxConsecutiveZero+1)/2 (maxConsecutiveZero为奇数数)