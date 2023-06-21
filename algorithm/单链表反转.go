package main

import "fmt"

type ListNode struct {
	data interface{}
	Next *ListNode
}

// 反转单链表 O(n)
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	i := 0
	for cur != nil {
		i++
		//获取下个节点temp
		temp := cur.Next
		//因要翻转，故cur.next=pre 下一个节点等于前一个节点
		cur.Next = pre
		//将当前节点赋值给前节点 pre后移
		pre = cur
		//将temp节点赋值给当前节点，用于继续往下执行
		cur = temp
		//平行赋值语法  可读性差
		//cur.Next,pre,cur=pre,cur,cur.Next
		PrintNode(string(rune(i)), pre)
	}

	return pre
}

func CreateNode(node *ListNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &ListNode{}
		cur.Next.data = i
		cur = cur.Next
	}
}

// 打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.data, " ")
	}
	fmt.Println()
}

func main() {
	var head = new(ListNode)
	CreateNode(head, 10)
	PrintNode("前：", head)
	yyy := reverseList(head)
	PrintNode("后：", yyy)

}
