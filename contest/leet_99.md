## Weekly Contest 99 解题报告

### 三位形体的表面积
[原题链接](https://leetcode-cn.com/contest/weekly-contest-99/problems/surface-area-of-3d-shapes/)

这道题乍一看挺复杂，因为两个相邻的坐标会有一面重叠。要确定一个立方体的某一面是否会重叠，需要知道它周围4个坐标的高度。如果坐标是一个随意大的数，这个就很不好处理了。不过看到数据范围，`N<=50`且坐标也是[0,50]，这样基本上就可以暴力枚举了。
* 首先对于某个坐标上的立方体，顶上和底下的面积是不会有重叠，所以可以直接+2。
* 对于立方体四个侧面，要看和它相邻的立方体的高度，假设当前立方体高为h，它左边紧邻的立方体高位v，如果h>v，则有h-v可以被计入表面积。剩下三个面同理。
* 需要注意处理边界case，即i=0 i=n-1和j=0 j=n-1的情况。

最终暴力枚举的复杂度是O(N^2), [AC代码](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_892/source.go)

### 特殊等价字符串
[原题链接](https://leetcode-cn.com/contest/weekly-contest-99/problems/groups-of-special-equivalent-strings/)

这道题读题读了很久很久很久才弄明白题意。其实就是让你把给定的字符串数组划分成N个小数组，每个小数组里面的字符串都可以`特殊等价`。根据题目对特殊等价的定义，我们可以想到一个非常简单的方式来判定两个字符串是不是特殊等价。

记录字符串偶数项每个字符出现的次数even[x]，再记录奇数项每个字符出现的次数odd[x]，如果两个字符串特殊等价，那么必然：
* S1.even[a..z] == S2.even[a..z]
* S1.odd[a..z] == S2.odd[a..z]

计算出一个字符串的even[a..z]和odd[a..z]，我们把even和odd称为该字符串的签名。因此可以枚举每个字符串，计算它的签名，然后看它的签名和哪个组的签名一致（一个组所有字符串签名一定是一致的）。如果没有组和它一样，则自成一组，不然则加入改组，最后返回一共有多少组即可。

复杂度O(N^3)，[AC代码](https://github.com/caibirdme/leetforfun/blob/master/leetcode/leet_893/source.go)

### 所有可能的完整二叉树
[原题链接](https://leetcode-cn.com/contest/weekly-contest-99/problems/all-possible-full-binary-trees/)

这道题比较考验大家设计递归的功力。如果我们用mem[N]来表示这道题要求的节点数为N的所有完整二叉树，根节点的左右子节点也是完整二叉树，假设左子树是节点数为x的完整二叉树，相应的，右子树肯定就是节点数为N-1-x的完整二叉树。即左子树有mem[x]种，右子树有mem[N-1-x]种，因此当根固定、左子树有x个节点，这样的完整二叉树一共有mem[x]*mem[N-1-x]种。每颗子树的构造就是个完全相同的子问题。
这道题看了代码挺简单，但是在设计递归的时候还是有点难度，比如根节点到底是在本次递归中生成还是在上次递归中生成然后传下去，还有很多别的问题需要考虑。不过这种构造类的递归方式，很值得一学。

AC代码:
```go
func allPossibleFBT(N int) []*TreeNode {
    hash := make(map[int][]*TreeNode)
    hash[1] = []*TreeNode{new(TreeNode),}
    return dfs(N, hash)
}

func dfs(N int, hash map[int][]*TreeNode) []*TreeNode {
    if v,ok := hash[N]; ok {
        return v
    }
    for i:=1; i<N; i+=2 {
        j := N-1-i
        left := dfs(i, hash)
        right := dfs(j, hash)
        for _,l := range left {
            for _,r := range right {
                root := new(TreeNode)
                root.Left = l
                root.Right = r
                hash[N] = append(hash[N], root)
            }
        }
    }
    return hash[N]
}
```

### 最大频率堆栈
[原题链接](https://leetcode-cn.com/contest/weekly-contest-99/problems/maximum-frequency-stack/)

这道题考察的是对数据结构的灵活运用。要求实现一个类似于栈的数据结构，push操作和普通栈一样，直接放到栈顶；但是pop操作不是弹出栈顶元素，而是弹出在栈中出现频次最多的元素（如果有多个就弹出离栈顶近的那个）。

我使用的是比较挫的办法，也比较直接，基本上就是硬刚。由于每次弹出要弹出现频次最大的元素（相同的话就是离栈顶最近的），因此我可以维护一个大顶堆，堆中每个元素保存某个数出现的频次以及该元素离栈顶最近的元素的下标。通过双关键字比较，就能维护这个大顶堆，每次pop时不是直接把堆顶元素删掉，而是把堆顶元素的frequency-1，然后调整它的位置。
为了通过元素值快速索引到它在堆中的位置，我们还需要一个map来保存 值x->堆中对应元素 的映射。

由于Go的map实现基于hashTable，因此Push和Pop的时间复杂都是常数

还有一种比较简洁的解决办法：
由于每次要弹出出现频次最高的元素，当频次一样才关心离栈顶的位置。因此可以对每个频次维护一个栈，即出现频次为1的数，放到栈1，出现频次为2的放到栈2，出现频次为k的放到栈k。假设栈k是目前下标最大的栈，意味只要从栈k的栈顶弹出元素即可。这种方法插入和删除都是O(1)，应该是最优雅的解法了。

## 总结
生成二叉树这道题让我学习了很多…构造类递归是常见的一类考题，要注意总结