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

	"github.com/afzalansari028/go-stack/mystack"
)

func main() {

	fmt.Println("let's play with stack using golang")

	stack := mystack.Stack{}

	stack.Push(10) // push an elemment into the stack int type
	stack.Push("a")
	stack.Display()
	stack.Pop() // remove an elememnt from the stack

	stack.Display()                            //display all elements
	fmt.Println("top element::", stack.Peek()) // see the top element

	fmt.Println("isempty::", stack.IsEmpty())
	fmt.Println("size::", stack.Size())
}

```
 Use Go command to run the code
```bash
$ go run main.go
```


