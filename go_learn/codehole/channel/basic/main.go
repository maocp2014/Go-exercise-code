package main

import "fmt"

/*
创建通道方法： make(type, size)  其中size参数是可选的

如果size不指定，则通道的容量为零，即非缓冲型通道，非缓冲型通道必须确保有协程正在尝试读取当前通道，否则写操作就会阻塞
直到有其它协程来从通道中读东西。非缓冲型通道总是处于既满又空的状态，与之对应的是有限定大小的通道，即缓冲型通道。
在 Go 语言里不存在无界通道，每个通道都是有限定最大容量的。

通道作为容器，它可以像切片一样，使用 cap() 和 len() 全局函数获得通道的容量和当前内部的元素个数。
通道一般作为不同的协程交流的媒介，在同一个协程里它也是可以使用的。

Go 语言为通道的读写设计了特殊的箭头语法糖 <-，让我们使用通道时非常方便。
把箭头写在通道变量的右边就是写通道，把箭头写在通道的左边就是读通道。一次只能读写一个元素。
*/

func main() {
	var ch chan int = make(chan int, 4) // 缓冲型通道，通道属于容器型变量
	for i := 0; i < cap(ch); i++ {
		ch <- i // 写通道，一次只能写一个元素
	}

	for len(ch) > 0 {
		value := <-ch // 读通道，一次只能读一个元素
		fmt.Println(value)
	}
}
