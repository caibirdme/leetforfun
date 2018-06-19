题目：

[原题链接](https://leetcode-cn.com/contest/weekly-contest-88/problems/shifting-letters/)

大意：

对'a'进行shift(3)操作结果为'd'
对'z'进行shift(1)操作结果为'a'

其实就是shift操作是循环的

然后给定一个字符串，每次都字符0..i进行shift[num[i]]操作，问最后得到的字符串是啥

题解：

这道题其实很简单，因为每次都对字符0..i操作，所以只需要倒着枚举shift，最后一个字符位移shift[last]，倒是第二个字符位移是shift[last-1]+shift[last]，以此类推。

我主要想玩点骚的，因为想起了以前做过的线段树题目，顺便再练练手，于是考虑用线段树+lazyload做。线段树用在这道题目简直就大材小用了，但是如果题目稍微改一下，每次对x..y进行shift(t)操作，这时候线段树就派上用场了。

我的代码是线段树的实现，但更简单的其实就是一个倒序枚举，10行就能搞定