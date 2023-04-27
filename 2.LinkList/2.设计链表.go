package main

/**
在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
*/

type MyLinkedList struct {
	isNil bool
	val   int
	next  *MyLinkedList
}

func Constructor() MyLinkedList {
	myLinkedList := new(MyLinkedList)
	myLinkedList.isNil = true
	return *myLinkedList
}

func (myList *MyLinkedList) Get(index int) int {
	list := myList
	for i := 0; i < index && list != nil; i++ {
		list = list.next
	}
	if list == nil || myList.isNil {
		return -1
	}
	return list.val
}

func (myList *MyLinkedList) AddAtHead(val int) {
	if !myList.isNil {
		node := &MyLinkedList{
			val:  myList.val,
			next: myList.next,
		}
		myList.val = val
		myList.next = node
	} else {
		myList.val = val
		myList.isNil = false
	}
}

func (myList *MyLinkedList) AddAtTail(val int) {
	if !myList.isNil {
		list := myList
		node := &MyLinkedList{
			val:  val,
			next: nil,
		}
		for list.next != nil {
			list = list.next
		}
		list.next = node
	} else {
		myList.val = val
		myList.isNil = false
	}
}

func (myList *MyLinkedList) AddAtIndex(index int, val int) {
	if !myList.isNil {
		list := myList
		if index <= 0 {
			myList.AddAtHead(val)
		} else {
			for i := 0; i < index && list != nil; i++ {
				list = list.next
			}
			if list == nil {
				myList.AddAtTail(val)
			} else {
				node := MyLinkedList{
					val:  list.val,
					next: list.next,
				}
				list.val = val
				list.next = &node
			}
		}
	} else {
		if index <= 0 {
			myList.val = val
			myList.isNil = false
		}
	}
}

func (myList *MyLinkedList) DeleteAtIndex(index int) {
	if !myList.isNil {
		head := &MyLinkedList{
			val:  0,
			next: myList,
		}
		for i := 0; i < index && head.next != nil; i++ {
			head = head.next
		}
		if head.next != nil {
			if head.next == myList {
				if myList.next != nil {
					myList.val = myList.next.val
					myList.next = myList.next.next
				} else {
					myList.val = -1
					myList.isNil = true
				}
			} else {
				head.next = head.next.next
			}
		}
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
