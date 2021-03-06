package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/03-Stacks-and-Queues/05-Array-Queue/arrayqueue"
)

func main() {
	queue := arrayqueue.New(20)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
