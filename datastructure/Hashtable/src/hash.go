package main

import "fmt"

// go hash表简单使用
func main() {
	fmt.Printf("hello world!\n")
	var m = make(map[string]int) //创建一个空白的hash
	m["str1"] = 10
	m["str2"] = 1000
	fmt.Printf("%d\n", m["str1"])
	// 遍历 hash
	for name, value := range m {
		fmt.Printf("%s\t%d\n", name, value)
	}
}
