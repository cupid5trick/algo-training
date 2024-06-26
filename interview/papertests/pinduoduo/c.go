package main
// 多多快递站共有n个快递点，n个快递点之间通过m条单向车道连接。快递员从任何一个快递站点出发，都无法通过单向车道回到该站点。
// 也就是说，n个快递点组成一张有向无环图。对于快递点u，如果对于所有的快递点 v(v!=u)， 快递员都可以从u走到v，或者从v走到u，那么则评定站点u为超级快递点。
// 请你帮忙算一算，一共有多少个超级快递点。
// 输入描述
// 第一行 2个数字n(2<=n<=3*10^5) , m(1<=m<=3*10^5) , n为快递点个数，m为单向车道个数。
// 接下来的m行每行两个数字 u,v(1<=u,v<=n, v!=u)，表示有一条站点u指向v的单向车道。
// 输出描述
// 请输出个数字，表示超级快递点的个数。