[题目链接](https://leetcode-cn.com/problems/word-ladder-ii/description/)

### 题解：

惭愧惭愧，这道题提交了好几次才AC，细节没处理好。其实算法挺简单的，就是一个简单的BFS，先把endWord入队，每发现一个可以转移到当前单词的单词，就把它入队，同时记录一下当前单词是由哪个单词扩展过来的，以便最后打印。如果扩展到了beginWord，就输出路径。里面比较多的细节需要处理。
