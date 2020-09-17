package main

import "fmt"

func main()  {
	//抱歉，运算符忘记在视频中讲了，但是没关系，我直接以代码输出给大家
	//算数运算符
	x := 1
	y := 2
	fmt.Println(x + y) //加法
	fmt.Println(x - y) //减法
	fmt.Println(x * y) //乘法
	fmt.Println(x / y) //除法
	fmt.Println(x % y) //取模，就是取余
	x++ //自增，相当于 x = x + 1
	x-- //自减，相当于 x = x - 1

	//关系运算符
	//==	检查两个值是否相等，如果相等返回 True 否则返回 False。	(A == B) 为 False
	//!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。	(A != B) 为 True
	//>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。	(A > B) 为 False
	//<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。	(A < B) 为 True
	//>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	(A >= B) 为 False
	//<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。	(A <= B) 为 True
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)

	//逻辑运算符
	//&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。	(A && B) 为 False
	//||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。	(A || B) 为 True
	//!	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。	!(A && B) 为 True
	a := true
	b := false
	fmt.Println( a && b)
	fmt.Println( a || b)
	fmt.Println( !a)

	//位运算符
	//p	q	p & q	p | q	p ^ q
	//0	0	0	0	0
	//0	1	0	1	1
	//1	1	1	1	0
	//1	0	0	1	1

	//&	按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。	(A & B) 结果为 12, 二进制为 0000 1100
	//|	按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或	(A | B) 结果为 61, 二进制为 0011 1101
	//^	按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。	(A ^ B) 结果为 49, 二进制为 0011 0001
	//<<	左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。	A << 2 结果为 240 ，二进制为 1111 0000
	//>>	右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。	A >> 2 结果为 15 ，二进制为 0000 1111

	fmt.Println( x & y)
	fmt.Println( x | y)
	fmt.Println( x ^ y)
	fmt.Println( x << y)
	fmt.Println( x >> y)
}