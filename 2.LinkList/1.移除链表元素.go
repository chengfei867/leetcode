package main

/**
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回新的头节点
*/

// ListNode 题目给的链表结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	iter := &ListNode{}
	iter.Next = head
	newHead = iter
	for iter.Next != nil {
		//如果当前节点的下一个节点是val,移除
		if iter.Next.Val == val {
			iter.Next = iter.Next.Next
		}
		if iter.Next != nil && iter.Next.Val != val {
			iter = iter.Next
		}
	}
	return newHead.Next
}
