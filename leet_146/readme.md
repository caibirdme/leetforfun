设计一个cache，支持以下操作：
* get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
* put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

要求插入和删除都是O(1)

题解：

其实也没什么算法，主要就是数据结构的实现。插入和删除O(1)显而易见就是用hashTable了。hashTable每个slot是一个链表，为了方便删除，链表最好是双向的。我这里直接用了go的标准库，因为hashTable和链表之类的之前的题目已经写过了，懒得写了。
LRU的关键是还需要一个链表来存储"每个key的使用情况“，其实就是：
* 每新来一个key就查到链表头
* SET一个已存在的key也把它放到链表头
* GET一个key就把它放到链表头
因此，链表尾的肯定就是一直都没有被GET和SET的。如果capacity满了，就把链表尾的删了，给新key腾空间