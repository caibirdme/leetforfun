题目：
判断字符串S是否可以由dict种元素组合而成，dict中每个元素可以使用多次。

S: "applepenapple"
dict: ["apple", "pen"]
=> true

S: "leetcode"
dict: ["leet", "code"]
=> true

S: "catsandog"
dict: ["cats", "dog", "sand", "and", "cat"]
=> false

S: "catsandsandcatsandcatdogdogcatcatsandcatscatdogsand"
dict: ["cats", "dog", "sand", "and", "cat"]
=> true

题解：

简单的动态规划，f[i]表示从S的第i个字符到末尾这个子串能否由dict构成。很显然，如果S[0..i-1]是dict中的一个字符串，那么问题成了判断子串S[i..last]是否可以由dict构成，递归一下就好了。判断S[0..i-1]是否是dict中的一个子串可以反过来，枚举dict，看dict[k]是不是S的前缀。