/*
* @Description:
* @Author: jafar du
* @Date: 2022/6/1 22:37
 */
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next // 把下一个节点缓存起来
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	//var pre *ListNode
	pre := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Printf("pre:%#v\n", pre)
	ret := reverseList(pre)
	fmt.Printf("ret:%#v\n", ret)
	for ret != nil {
		fmt.Print(ret.Val, "->")
		ret = ret.Next
	}
}
