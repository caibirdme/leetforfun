题目：
> Your car starts at position 0 and speed +1 on an infinite number line.  (Your car can go into negative positions.)
> Your car drives automatically according to a sequence of instructions A (accelerate) and R (reverse).
> When you get an instruction "A", your car does the following: position += speed, speed *= 2.
> When you get an instruction "R", your car does the following: if your speed is positive then speed = -1 , otherwise speed = 1.  (Your position stays the same.)
> For example, after commands "AAR", your car goes to positions 0->1->3->3, and your speed goes to 1->2->4->-1.
> Now for some target position, say the length of the shortest sequence of instructions to get there.

这道题还是有点意思的，一开始觉得像DP，但又觉得速度是前后关联的，似乎和DP的“无后效性”相悖？没有深想于是换了个思路。
非常容易想到的就是BFS，因为每次要么加速要么转向，于是乎每次可以新产生两个状态加入队列，然后再不断派生状态，直到第一次走到终点。
BFS需要注意的就是，要避免一个状态反复出现。比如“以速度16通过位置20”，如果之前这个状态已经出现过，之后再走到这个状态就可以剪枝了。
另一个比较重要的剪枝是`2倍target`剪枝。啥意思呢？ 假设target是100，2倍target就是200。想象一下，你的目标是要走到100，但是你现在已经走到200了。那么如果你接着向前走，只会越来越远。如果你调头，需要一个R指令，然后问题变成了`初始速度是-1，从200走到100最少需要多少个指令`，这本质上和`初始速度+1从0走到100`是一个问题。也就是说，假设从0走到100的最少步数是y，如果你花了x步走到2倍target处，这时候你倒回到100处，总共花的步数x+y，显然x+y>x，所以如果走到2倍target处，就可以把那个状态剪掉了。大于2倍target的就更好想了，必然要剪掉。你会问，大于2倍target的地方再掉头往回走，会不会刚好连续几个加速就能到（也就是距离刚好是2^n-1）。这个其实也很好想，如果在target+2^n-1的地方掉头，确实可以n步走回target，但是很显然，你完全可以在target+2^(n-1)-1处就掉头，这样比在target+2^n-1时再掉头要少一步。
综上，2倍target剪枝意思说，如果走到pos >= target*2 了，肯定就不是最优的走法了，剪掉。

提交了之后，发现大家有用DP做的，仔细看了下，确实DP有更简单的做法。
设DP[i]表示以初始速度1从0走到i需要的最少步数。
* 首先可以想到，如果i == 2^n-1，那么DP[i]必然就是n了，也就是说连续几个“AAA……”指令刚好走过去。
* 如果 2^(n-1)-1 < i < 2^n-1，也就是说，连续几个AAA指令之后，离i还有点距离，但是再来一个A又走过了。
    * 走过了再掉头，掉头之后速度是-1，这时候往回走，就是原问题的一个子问题。所以需要的步数是 n个A + 1个R掉头 + DP[2^n-1 - target]
    * 如果离i还差点距离，还可以连续两个掉头，把速度重新变回1，然后剩下的距离又是个相同子问题了。这种走法需要的步数是 dp[2^(n-1)-1] + 2个R掉头 + dp[target-(2^(n-1)-1)]
    * 如果离i还差点距离，也可以提前掉头往后挪点，如果提前掉头往后挪到位置k，且k~i的距离刚好是2^t-1，岂不是美滋滋？到底挪多少呢？这个需要枚举，假设往回走m步，那么总的步数就是 n-1个A + 1个R掉头 + m步往回走 + 1个R掉头 + dp[i-k]