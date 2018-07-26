之前看有同学建议大家一起没事儿做做leetcode，我最近也重新开始刷刷题，所以之后如果没有特别的情况，我都会把每周contest的解题报告发到这里。如果不知道leetcode每周contest的同学可以看[这里](https://leetcode-cn.com/contest/)，题目难度适中

# contest 94

## 叶子相似的树
[原题链接](https://leetcode-cn.com/contest/weekly-contest-94/problems/leaf-similar-trees/)

考虑一个二叉树的所有叶子。这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
![tree](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/16/tree.png)
举个例子，给定一个如上图所示的树，其叶值序列为 (6, 7, 4, 9, 8)
如果两个二叉树的叶值序列相同，我们就认为它们是 叶相似的。
如果给定的两个头结点分别为 root1 和 root2 的树是叶相似的，返回 true；否则返回 false 。

### 题解
因为叶子节点是从左到右，通过一次后序遍历就能得到，每次到叶子节点就把值加入数组。因此这道题其实就是在考后序遍历，送分题
AC代码: [872](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_872/source.go)

## 行走机器人模拟
[原题链接](https://leetcode-cn.com/contest/weekly-contest-94/problems/walking-robot-simulation/)

题目其实就是说，有三种指令：
* 前进K步（K<=9）
* 左转
* 右转

然后地图上有障碍，如果机器人前进的过程中遇到障碍，那就停在障碍前面等待下一个指令。
机器人一开始在原点(0,0)，朝北，输入一串指令，问机器人最远离原点有多远

### 题解
这其实也是一道简单题，由于K<=9，甚至可以一步一步地模拟看有没有障碍。如果不加K<=9，这道题会复杂一点。这道题我由于一开始读题有问题，以为是问所有指令执行完离原点有多远，因此没有在每执行完一个指令都算一下距离，导致一直错…尴尬。其实很简单的题目，就不多说了。
AC代码：[873](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_874/source.go)

## 爱吃香蕉的珂珂
[原题链接](https://leetcode-cn.com/contest/weekly-contest-94/problems/koko-eating-bananas/)

珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。
珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。  
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

### 题解
简化一下题目：一共有N个数，科科每次操作可以选一个数减去K，要求H次操作后要把所有数变成非正数，问这个K最小是多少。
* 因为有N个数，每次操作一个，因此如果H<N，肯定是无解的
* 如果N等于H，就代表说每次操作都得把那个数变成<=0，因此这个K一定要和N个数中最大的那个数相同，才能保证在对那个数执行操作时把它置零

也就是说，K最大不超过N个数的最大值。
其实这种`最大值最小`的问题，基本上都可以往二分上面想。
如果K取M时，H次操作可以把N个数置零(<=0)，那么显而易见的，K>=M时，同样满足。因此我们可以二分一个答案然后验证。
显而易见的，1<= K <= maxN，所以我们可以先尝试看 k= (1+maxN)/2
当k = mid = (1+maxN)/2时，看能否通过H次操作完成"置零"，如果可以，那么答案一定<=mid，如果不行，答案一定>mid。其实也是简单题，就不多说了。
[AC代码](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_875/source.go)

## 最长的斐波那契子序列的长度
[原题链接](https://leetcode-cn.com/contest/weekly-contest-94/problems/length-of-longest-fibonacci-subsequence/)

如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
* n >= 3
* 对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}

给定一个严格递增的正整数数组形成序列，找到 A 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。

（回想一下，子序列是从原序列 A 中派生出来的，它从 A 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）

### 题解
根据题目定义, fib数列中x[i]=x[i-1]+x[i-2]，这意味着，只要知道了fib数列的前两项，我们就可以推出后面的无数项，回想一下我们熟知的fib数列 0 1 1 2 3 5 8 .... 只要给你 0 1，后面的就很容易算出来了。
因此一个简单的想法就是，枚举fib数列的前两项，然后基于这两个数看能扩展出多长的序列。比如已知前两项是a b，第三项就是a+b。a+b必须在原序列A中，a b a+b才是A的子序列。怎么判断a+b是否在A中呢？显然，二分查找，因为A是严格递增的。复杂度是O(N^2*logN)。
[AC代码](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_873/source.go)

## 总结
这场比赛其实还是比较简单的，后两题也不过是medium。本场算是个二分查找专项训练吧

