package main

import (
	"fmt"
	"strings"
)

func main() {

	//m := make(map[int][]int)
	//fmt.Println(howSum(7, []int{2, 3, 5, 7}, m))
	//fmt.Println(bestSum(100, []int{1, 4, 5, 25}, m))

	// ms := make(map[string]bool)
	// fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, ms))
	// mi := make(map[string]int)
	// fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}, mi))
	fmt.Println(allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))

}

// find any one combination of making a target sum from the given nums array
func howSum(target int, nums []int, memo map[int][]int) []int {

	/*
		target = 7, nums= [2,3,5,7]
		output combination {{7}, {5,2}, {3,2,2}}
			i = 0
			reminder -  7-2=5, 7-3=4, 7-5=2, 7-7=0
	*/

	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target < 0 {
		return nil
	}

	if target == 0 {
		return []int{}
	}

	for i := 0; i < len(nums); i++ {
		reminder := target - nums[i]
		reminderResult := howSum(reminder, nums, memo)

		if reminderResult != nil {
			reminderResult = append(reminderResult, nums[i])
			memo[target] = reminderResult
			return memo[target]
		}
	}

	memo[target] = nil
	return nil
}

// find the best combination of making a target sum from the given nums array
func bestSum(target int, nums []int, memo map[int][]int) []int {

	/*
		target = 7, nums= [2,3,5,7]
		 {{7}, {5,2}, {3,2,2}}
		output best combination {7}
			i = 0
			reminder -  7-2=5, 7-3=4, 7-5=2, 7-7=0
	*/

	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if target < 0 {
		return nil
	}

	if target == 0 {
		return []int{}
	}

	var smallcomb []int
	for i := 0; i < len(nums); i++ {
		reminder := target - nums[i]
		reminderCombination := bestSum(reminder, nums, memo)

		if reminderCombination != nil {
			n := nums[i]
			comb := append([]int{}, reminderCombination...)
			comb = append(comb, n)
			if len(smallcomb) == 0 || len(smallcomb) > len(comb) {
				smallcomb = comb
			}
		}
	}

	memo[target] = smallcomb
	return smallcomb
}

func canConstruct(target string, wordBank []string, memo map[string]bool) bool {

	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if len(target) == 0 {
		return true
	}

	for _, v := range wordBank {
		if strings.Index(target, v) == 0 {
			fmt.Println("prefix words --- ", v)
			remainder := strings.Trim(target, v)
			if canConstruct(remainder, wordBank, memo) == true {
				memo[target] = true
				return true
			}

		}
	}

	memo[target] = false
	return false
}

func countConstruct(target string, wordBank []string, memo map[string]int) int {
	// purple , {purp, p, ur, le, purpl}
	count := 0

	if _, ok := memo[target]; ok {
		return memo[target]
	}
	if len(target) == 0 {
		return 1
	}

	for _, v := range wordBank {

		if strings.Index(target, v) == 0 {
			fmt.Println("prefix are ---", v)
			reminder := strings.Trim(target, v)
			tempCount := countConstruct(reminder, wordBank, memo)
			count = count + tempCount
			memo[target] = count

		}
	}

	count = memo[target]
	return count
}

func allConstruct(target string, wordBank []string) [][]string {

	if target == "" {
		return [][]string{{}}
	}
	var result [][]string

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := strings.Trim(target, word)

			allsuffix := allConstruct(suffix, wordBank)
			//fmt.Println("allreminders----", allreminders) // this fmt.println gives empty 2D string,so no point
			for _, eachreminder := range allsuffix {
				result = append(result, append([]string{word}, eachreminder...))
			}

		}
	}

	return result
}
