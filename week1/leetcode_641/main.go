package main

type MyCircularDeque struct {
	Size  int
	Queue []int
	Head  int
	Tail  int
}

func Constructor(k int) MyCircularDeque {
	size := k + 1
	var queue = make([]int, size)
	return MyCircularDeque{Size: size, Queue: queue, Head: 0, Tail: 0}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.Head = (this.Head - 1 + this.Size) % this.Size
		this.Queue[this.Head] = value
		return true
	}
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.Queue[this.Tail] = value
		this.Tail = (this.Tail + 1) % this.Size
		return true
	}
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.Head = (this.Head + 1) % this.Size
		return true
	}
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.Tail = (this.Tail - 1 + this.Size) % this.Size
		return true
	}
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.Queue[this.Head]
	}
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.Queue[(this.Tail-1+this.Size)%this.Size]
	}
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Head == this.Tail
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.Tail+1)%this.Size == this.Head
}

func main() {
	circularDeque := Constructor(3)

	circularDeque.InsertLast(1)  // 返回 true
	circularDeque.InsertLast(2)  // 返回 true
	circularDeque.InsertFront(3) // 返回 true
	circularDeque.InsertFront(4) // 已经满了，返回 false
	circularDeque.GetRear()      // 返回 2
	circularDeque.IsFull()       // 返回 true
	circularDeque.DeleteLast()   // 返回 true
	circularDeque.InsertFront(4) // 返回 true
	circularDeque.GetFront()     // 返回 4

}
