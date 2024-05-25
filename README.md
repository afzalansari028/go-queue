**Queue Data Structure Golang**
====================================
Queue is a linear data structure which follows the FIFO(First In First Out) policy. You can remove an element from front side and insert from rear side.

**Get started**
===================
## Installation

With the help of below Go command, simply install the `myqueue` package
```bash
 $ go get github.com/afzalansari028/go-queue/myqueue
```
## Example:
```bash
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

	queue.Dequeue()

	fmt.Println("Peek::", queue.Peek())

	queue.Dispay()

}

```
 Use Go command to run the code
```bash
$ go run main.go
```


