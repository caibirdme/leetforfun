### 题解

设：i个不同的数一共能组成f(i)棵BST

有一个显而易见的事实是，BST的个数只和i有关，和这i个数本身无关。比如123和345，能组成的BST的个数是一样的。

我们可以枚举root，假设root是i，那么它的左子树就是1..i-1共i-1个节点，右子树是i+1..n共n-i个节点。因此当root是i时，一共有f[i-1]*f[n-i]种不同的BST。