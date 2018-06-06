实现一个函数，其功能是支持带 .* 的正则匹配。

不要想太多什么自动机之类的，其实就是一个递归下降。

需要分情况讨论：

先考虑pattern的长度，当且仅当pattern长度大于1才有可能出现*
* 如果pattern为空，s必须为空
* 如果pattern长度是1，s长度也必须是1且s==pattern或pattern=='.'
* 如果pattern长度大于等于2：
    * 如果pattern第一个字符后面跟的不是*, return equal(s[0], pattern[0]) && isMatch(s[1:], pattern[1:])
    * 如果pattern第一个字符后面是*，则可以匹配0个或多个：
        * 匹配0个就是 isMatch(s, pattern[2:])
        * 匹配多个就是 equal(s[0], pattern[0]) && isMatch(s[1:], pattern)。这里可能有人会有点疑惑，为什么匹配多个是isMatch(s[1:], pattern)。仔细想想递归的过程就明白了。