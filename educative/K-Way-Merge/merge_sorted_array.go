package main

import (
	"fmt"
	"strings"
)

/*

1. Initialize two pointers p1 and p2 that point to the last data elements in nums1 and nums2 respectively.
2. Initialize a pointer p, that points to the last element of nums1
3. If the value at p1 is greater than the value of p2, set the value at p equal to p1 and decrement p1 and p by 1
4. Else , if the value at p2 is greater than the value at p1, set the value at p equal to p2 and decrement p2 by 1
5. Continue the traversal until nums2 is merged with nums1

The time complexity is O(n+m)O(n+m), where nn and mm are the counts of initialized elements in the two arrays.
The space complexity is O(1)O(1) because we only use the space required for three indices.

*/

func mergeSorted(nums1 []int, m int, nums2 []int, n int) []int {

	// Replace this placeholder return statement with your code
	p1 := m - 1
	p2 := n - 1

	for p := m + n - 1; p >= 0; p-- {
		if p2 < 0 {
			break
		}

		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1 -= 1
		} else {
			nums1[p] = nums2[p2]
			p2 -= 1
		}
	}

	return nums1
}

func main() {
	m := []int{9, 2, 3, 1, 8}
	n := []int{6, 1, 4, 2, 1}
	nums1 := [][]int{
		{23, 33, 35, 41, 44, 47, 56, 91, 105, 0, 0, 0, 0, 0, 0},
		{1, 2, 0},
		{1, 1, 1, 0, 0, 0, 0},
		{6, 0, 0},
		{12, 34, 45, 56, 67, 78, 89, 99, 0},
	}
	nums2 := [][]int{
		{32, 49, 50, 51, 61, 99},
		{7},
		{1, 2, 3, 4},
		{-99, -45},
		{100},
	}

	for index, value := range m {
		fmt.Printf("%d.\tnums1: %s, m: %d\n", index+1, strings.Replace(fmt.Sprint(nums1[index]), " ", ", ", -1), value)
		fmt.Printf("\tnums2: %s, n: %d\n", strings.Replace(fmt.Sprint(nums2[index]), " ", ", ", -1), n[index])
		result := mergeSorted(nums1[index], value, nums2[index], n[index])
		fmt.Printf("\tMerged list: %s\n", strings.Replace(fmt.Sprint(result), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
