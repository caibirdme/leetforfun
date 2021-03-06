题意：

`*` 可以匹配任意长字符串（包括空串）
`?` 可以匹配任何一个字符

写一个函数判断pattern和待匹配字符串是否匹配

题解：

这题和正则通配很像，就是一个小改动。但是直接裸做超时了……

实际上优化也非常容易想到，就是一个简单的动态规划：

mem[i][j]表示 p[0..i]和s[0..j]是否匹配，则有mem[i][j]= ：

* mem[i-1][j-1] (p[i] == s[j] || p[i] == '?') 如果两个串最后一个字符能匹配
* mem[i][j-1] || mem[i-1][j]  (p[i] == '*') 如果p串最后一个字符为*

还有些小优化可以减少搜索量：
* 连续的`*`可以合并成一个`*`，如p="ab****\*c" 和p="ab*c"是等价的
* 如果p串形如: "asadasdasqowie??asdase\*asddd"，在p中第一个*号之前的串可以先和S串进行逐字匹配
---

我采用了记忆化搜索，但是犯了一个非常愚蠢的错误，连续超时了好几次才发现问题……
我定义存储搜索结果的二维表定义成了 `var mem [][]bool`，因此代码中当且仅当mem[i][j] == true才直接返回，因为false是mem的默认值，没法判断到底是mem[i][j]没搜索过还是搜索结果为false。 事实上false和unset不是相同的状态……所以最后把mem的定义换成了：

```go
var mem [][]byte

const (
    _ byte = iota
    yes
    no
)

//...
if mem[i][j] == yes {
    return true
} else if mem[i][j] == no {
    return false
}
```
