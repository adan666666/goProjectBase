package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //只定义无需赋值

/*
*
两个协程，操作同一个管道  --进行通讯
*/
func main() { //主线程
	//写协程和读协程共同操作同一个管道-》定义管道：
	intChan := make(chan int, 50)
	wg.Add(2)
	//开启读和写的协程：
	go writeData(intChan)
	go readData(intChan)
	//主线程一直在阻塞，什么时候wg减为0了，就停止
	wg.Wait()
}

// 写：
func writeData(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入的数据为：", i)
		time.Sleep(time.Second)
	}
	//管道关闭：
	close(intChan)
}

// 读：
func readData(intChan chan int) {
	defer wg.Done()
	//遍历：
	for v := range intChan {
		fmt.Println("读取的数据为：", v)
		time.Sleep(time.Second)
	}
}
