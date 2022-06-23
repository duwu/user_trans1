package main

import (
	"container/list"
	"fmt"
)

func NewQueue() {
	// 生成队列
	l := list.New()

	// 入队
	l.PushBack(1)
	l.PushBack(2)
	// 出队
	i1 := l.Front()
	l.Remove(i1)
	fmt.Printf("%d\n", i1.Value)
	i2 := l.Front()
	l.Remove(i2)
	fmt.Printf("%d\n", i2.Value)

	//入栈
	l.PushBack(1)
	l.PushBack(2)
	// 出栈
	i3 := l.Back()
	l.Remove(i3)
	fmt.Printf("%d\n", i3.Value)
	i4 := l.Back()
	l.Remove(i4)
}
