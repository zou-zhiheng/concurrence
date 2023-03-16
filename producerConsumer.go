package main

import (
	"fmt"
	"time"
)

//生产者：生成factor的整数倍的序列
func producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

//消费者
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func ProducerConsumerDemo() {

	ch := make(chan int, 64) //成果序列

	go producer(3, ch) //生成3的倍数的序列
	go producer(5, ch) //生成5的倍数的序列
	go consumer(ch)    //消费生成的序列
	time.Sleep(5 * time.Second)

}
