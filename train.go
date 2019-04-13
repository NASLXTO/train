package main

import (
	"fmt"
	"time"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if p, ok := m[target-v]; ok {
			return []int{p, i}
		}
		m[v] = i
	}
	return []int{}
}

//ListNode 链表.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	i := 10
	j := 0
	for l1 != nil || l2 != nil || j != 0 {
		var sum int
		var tmp ListNode

		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		sum += j

		if l1 == nil && l2 == nil && j != 0 {
			tmp = ListNode{j, nil}
			j = 0
		} else {
			tmp = ListNode{sum % i, nil}
			j = sum / i
		}
		if l3 == nil {
			l3 = &tmp
		} else {
			tmp.Next = l3
			l3 = &tmp
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	var l4 *ListNode
	for l3 != nil {
		tmp := ListNode{l3.Val, nil}
		if l4 == nil {
			l4 = &tmp
		} else {
			tmp.Next = l4
			l4 = &tmp
		}
		l3 = l3.Next
	}
	return l4
}

func countsub(n int, s string) int {
	if s == "" {
		return n
	} else {
		m := make(map[int32]int)
		j := 0
		for i, v := range s {
			if p, ok := m[v]; ok {
				if i-p > n {
					n = i - p
				}
				fmt.Println(n, p, i)
				return countsub(n, s[p+1:])
			}
			m[v] = i
			j++
		}
		if j >= n {
			return j
		} else {
			return n
		}
	}
}

func lengthOfLongestSubstring(s string) int {
	//	max := 0
	//	max = countsub(max, s)
	//	return max
	return countsub(0, s)

}

func main() {
	t := time.Now()

	//1.
	//	result := twoSum([]int{2, 7, 11, 15}, 9)
	//	fmt.Println(result)

	//2.
	//	result := addTwoNumbers(&ListNode{9, &ListNode{9, nil}}, &ListNode{1, nil})
	//	for result != nil {
	//		fmt.Println(result)
	//		result = result.Next
	//	}

	//3.
	//aaaaa abcdef abcdaef abcabcb cddb
	result := lengthOfLongestSubstring("cdd")
	fmt.Println(result)

	fmt.Println("Timing:", time.Since(t))

}
