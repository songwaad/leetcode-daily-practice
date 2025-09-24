// 4. Median of Two Sorted Arrays
// https://leetcode.com/problems/median-of-two-sorted-arrays/description/?envType=problem-list-v2&envId=array
// Time Complexity: O(nlogn)

package arrayandstring

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	rs := append(nums1, nums2...)
	sort.Ints(rs)

	n := len(rs)
	if n%2 == 0 {
		return float64(rs[n/2]+rs[n/2-1]) / 2
	} else {
		return float64(rs[n/2])
	}

}
