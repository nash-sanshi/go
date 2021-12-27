package main

import "fmt"

func main() {
	//定义数组
	var myArray = [5]string{"I", "am", "stupid", "and", "weak"}
	//普通遍历
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
	//range遍历
	for index, value := range myArray {
		fmt.Println(index, value)
	}
	//通过普通遍历修改值
	for i := 0; i < len(myArray); i++ {
		myArray[2] = "smart"
		myArray[4] = "strong"
		fmt.Println(myArray[i])
	}
	//通过index修改值，生效
	for index, value := range myArray {
		if index == 2 {
			myArray[index] = "smart"
		}
		if index == 4 {
			myArray[index] = "strong"
		}
		fmt.Println(index, value)
	}
	//通过value修改值，生效
	for index, value := range myArray {
		if value == "stupid" {
			value = "smart"
		}
		if value == "weak" {
			value = "strong"
		}
		fmt.Println(index, value)
	}
}
