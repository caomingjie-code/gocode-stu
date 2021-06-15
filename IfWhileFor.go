package main

/**
 * 学习if else while for 语法
 */
func main() {
	ifCode() //if 代码编写
}

func ifCode() {
	var num = 1
	//第一种书写格式
	if num >= 1 {
		println("if (num >= 1) {}")
	}
	//第二种书写格式
	if num >= 1 {
		println("if num >= 1 {}")
	}

	//if (num=2; num>=1){} 不可以这么写
	//第三种书写格式
	if num = 2; num >= 2 {
		println("if num =2 ; num>=2 {}")
	}
}
