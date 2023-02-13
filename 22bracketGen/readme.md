[22. 括号生成](https://leetcode.cn/problems/generate-parentheses/)
重点:
回溯法: 利用递归+剪枝,实现比暴力法更高效的解题方式
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
提示：

1 <= n <= 8
Related Topics
字符串
动态规划
回溯


题解1 
暴力法,生成所有题解, 之后利用函数筛选有效题解

题解2 
回溯法.
思路: 利用递归生成所有可能情况,之后利用剪枝法去掉错误的生成路径

推荐题解
[有go题解和手画有效剪枝图](https://leetcode.cn/problems/generate-parentheses/solution/shou-hua-tu-jie-gua-hao-sheng-cheng-hui-su-suan-fa/)
[带满二叉树图](https://leetcode.cn/problems/generate-parentheses/solution/sui-ran-bu-shi-zui-xiu-de-dan-zhi-shao-n-0yt3/)