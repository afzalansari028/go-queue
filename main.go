package main

import (
	"fmt"

	"github.com/afzalansari028/go-queue/myqueue"
)

func main() {
	fmt.Println("queueu!!!!!!")

	queue := myqueue.Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)
	queue.Enqueue(10)

	fmt.Println("Peek::", queue.Peek())

	queue.Dispay()

}
