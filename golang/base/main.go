package main

import (
	"fmt"
	"sort"
)

// p判断一个元素在数组中只出现一次
func findSingleNumber(nums []int) int {
	// 创建一个map来存储每个数字的出现次数
	countMap := make(map[int]int)

	// 第一次遍历：统计每个数字的出现次数
	for _, num := range nums {
		fmt.Println(num)
		countMap[num]++
	}

	// 第二次遍历：找出只出现一次的数字
	for num, count := range countMap {
		fmt.Println(num)
		if count == 1 {
			return num
		}
	}

	// 如果没有找到，返回-1（根据题目描述，这种情况不应该发生）
	return -1
}

// 判断是否为回文数
func isPalindrome(x int) bool {
	// 特殊情况处理：
	// 如果x为负数，不是回文数
	// 如果x最后一位是0，只有当x为0时才是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为偶数时，x == revertedNumber
	// 当数字长度为奇数时，x == revertedNumber / 10
	return x == revertedNumber || x == revertedNumber/10
}

// 有效的括号
func isValid(s string) bool {
	// 使用切片模拟栈
	stack := []rune{}

	// 创建映射表，用于匹配括号
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是右括号
		if matchingLeft, isRight := bracketMap[char]; isRight {
			// 检查栈是否为空或栈顶元素不匹配
			if len(stack) == 0 || stack[len(stack)-1] != matchingLeft {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈中
			stack = append(stack, char)
		}
	}

	// 如果栈为空，说明所有括号都匹配成功
	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串为基准
	for i := 0; i < len(strs[0]); i++ {
		// 遍历后续字符串
		for j := 1; j < len(strs); j++ {
			// 如果当前索引超过某个字符串的长度，或者字符不匹配
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 整数数组 +1
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按区间的起始值排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片，包含第一个区间
	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		// 如果当前区间的起始值小于等于上一个区间的结束值，则合并
		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 否则，添加当前区间
			merged = append(merged, current)
		}
	}

	return merged
}

// 两数之和
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, exists := numMap[complement]; exists {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	// 测试用例
	/*	nums := []int{4, 1, 2, 1, 2}
		result := findSingleNumber(nums)
		fmt.Printf("数组 %v 中只出现一次的元素是: %d\n", nums, result)

		// 另一个测试用例
		nums2 := []int{2, 2, 1}
		result2 := findSingleNumber(nums2)
		fmt.Printf("数组 %v 中只出现一次的元素是: %d\n", nums2, result2)

		// 测试用例
		testCases := []int{121, -121, 10, 12321, 0, 1221}
		for _, num := range testCases {
			fmt.Printf("%d 是回文数吗？ %t\n", num, isPalindrome(num))
		}*/
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""

	fmt.Println(plusOne([]int{2, 3, 4})) // ""
	fmt.Println(plusOne([]int{0, 9, 9})) // ""

	fmt.Println(removeDuplicates([]int{1, 2, 2})) // 删除数组中的重复元素  数组长度变成2

	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // 输出: [[1 6] [8 10] [15 18]]

	nums1 := []int{2, 3, 4} //  target 7
	fmt.Println(twoSum(nums1, 7))

}
