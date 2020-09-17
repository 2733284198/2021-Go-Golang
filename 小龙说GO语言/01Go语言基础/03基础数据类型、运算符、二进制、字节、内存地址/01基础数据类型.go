package main
//Main file has non-main package or doesn't contain main function
//你如果想要run一个文件的话，那么你必须要有main包和main主函数呀

import "fmt"

func main (){
	//fmt.Println("Hello World")
	//整型 int int8 int16 int32 int64
	//go编译器会默认把这一行话翻译成 var xiaolong int = 0
	var xiaolong int
	xiaolong = 2021
	fmt.Println(xiaolong)

	var fudian float32
	fudian = 1.28
	fmt.Println(fudian)

	//go语言编译器会把这段话翻译成：var xiaolong1 int = 2021
	xiaolong1 := 2021 //节约生命,主流的用法
	fmt.Println(xiaolong1)

	//有符号无符号
	var wufuhao uint
	fmt.Println(wufuhao)

	//复数类型
	//字符串
	var zifuchuan string
	zifuchuan = "xiaolongtalkgo2021"
	fmt.Println(zifuchuan)

	zifuchuan1 := "我是节约生命的字符串申明, 因为有了GO语言编辑器帮我自动做了强类型"
	fmt.Println(zifuchuan1)

	//布尔类型
	var xiaolongx bool //true || false
	xiaolongx = true
	fmt.Println(xiaolongx)
	//越抽象才越接近本质，越本质才越细节

	//字符类型
	var zifu byte
	zifu = 'a'
	fmt.Println(zifu)


}