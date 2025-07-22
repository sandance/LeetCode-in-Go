package main

import (
	"fmt"
	"strings"
)

/*
Algo:
1. Traverse in nums using the fast and slow pointers.
2. Move pointers until they meet. The slow pointer moves once and the fast pointer moves twice as fast as the slow pointer.
3. After the pointers meet, traverse in nums again.
4. Move the slow pointer from the start of nums and the fast pointer from the meeting point at the same speed (one step) until they meet again.
5. Return the fast pointer.
*/

func findDuplicate(nums []int) int {
	fast, slow := nums[0], nums[0]

	// PART #1
	// Traverse in array until the intersection point is found
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		// Break the loop when slow pointer becomes equal to the fast pointer, i.e.,
		// if the intersection is found
		if fast == slow {
			break
		}
	}

	// PART #2
	// Make the slow pointer point the starting position of an array again, i.e.,
	// start the slow pointer from starting position
	slow = nums[slow]
	for slow != fast {
		slow = nums[slow]
		// Move the fast pointer slower than before, i.e., move the fast pointer
		// using the nums[fast] flow
		fast = nums[fast]
	}
	return fast
}

func main() {
	nums := [][]int{
		{1, 3, 2, 3, 5, 4},
		{2, 4, 5, 4, 1, 3},
		{1, 6, 3, 5, 1, 2, 7, 4},
		{1, 2, 2, 4, 3},
		{3, 1, 3, 5, 6, 4, 2},
	}

	for i, num := range nums {
		fmt.Printf("%d.\tnums = %s\n", i+1, strings.Replace(fmt.Sprint(num), " ", ", ", -1))
		fmt.Printf("\tDuplicate number = %d\n", findDuplicate(num))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
