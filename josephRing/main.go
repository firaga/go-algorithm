package josephRing

//n 人数
//m 淘汰第m个人
//ret 幸存者在最开始的下标值
//https://blog.csdn.net/u011500062/article/details/72855826
//https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	return (lastRemaining(n-1, m) + m) % n
}

/*
定义函数f(n,m)为当前的幸存者下标.
则f(n-1,m) = f(n,m)-m
因为递归无法从上往下算,
且f(1,m) = 0
且有f(n,m) = (f(n-1,m)+m)%n(因为有环,取模)
可以递归从下向上计算出f(n,m)
转为算法实现
*/

/*
re.md为从上往下推,更合理一些
*/
