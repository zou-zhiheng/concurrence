package  main

import (
	"concurrence/ProducerConsumer"
	"concurrence/WorkPool"
	"fmt"
)

func main(){
	fmt.Println("start")
	ProducerConsumer.ProducerConsumerDemo()
	WorkPool.Demo()
}