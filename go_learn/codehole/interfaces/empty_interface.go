package main

import (
	"fmt"
)

func empty_interface1() {
	// 类型推导以及赋值
	// map的值为interface{}，表示value可以为任意类型，但是map的value类型是interface{}类型
	var user = map[string]interface{}{
		"age":     30,
		"address": "Beijing Tongzhou",
		"married": true,
	}

	fmt.Println(user)
	// 类型断言，接口类型转换成目标类型
	// user字典中变量的类型是map[string]interface{}，字典中得到的value类型是interface{}，需要通过类型转换才能得到期望的变量
	var age = user["age"].(int)
	var address = user["address"].(string)
	var married = user["married"].(bool)

	fmt.Println(age, address, married)
}

func empty_interface2() {
	/*
		# []T不能直接赋值给[]interface{}，它们是不同的类型
		t := []int{1, 2, 3, 4}
		var s []interface{} = t  // cannot use t (type []int) as type []interface {} in assignment
	*/

	// 正确做法
	t := []int{1, 2, 3, 4}
	s := make([]interface{}, len(t))

	//s = append(s, t...)  // cannot use t (type []int) as type []interface {} in append
	for i, v := range t {
		// fmt.Println(v)
		// s = append(s, v)
		s[i] = v
	}
	fmt.Println(s)
}

func main() {
	empty_interface1()
	empty_interface2()
}
