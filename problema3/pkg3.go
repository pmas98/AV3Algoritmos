package problema3

import (
	"fmt"
)

func findSubsetsUtil(nums []int, index int, currentSubset []int, currentSum int) bool {
	if currentSum == 0 && len(currentSubset) > 0 {
		fmt.Println("Subset found:", currentSubset)
		return true
	}
	print("currentSubset: ", fmt.Sprint(currentSubset), " currentSum: ", currentSum, "\n")
	print("index: ", index, " len(nums): ", len(nums), "\n")
	for i := index; i < len(nums); i++ {
		currentSubset = append(currentSubset, nums[i])

		if findSubsetsUtil(nums, i+1, currentSubset, currentSum+nums[i]) {
			return true
		}

		currentSubset = currentSubset[:len(currentSubset)-1]
		print("currentSubset: ", fmt.Sprint(currentSubset), " currentSum: ", currentSum, "\n")

	}

	return false
}

func findSubsets(nums []int) {
	currentSubset := []int{}
	if !findSubsetsUtil(nums, 0, currentSubset, 0) {
		fmt.Println("No subset with sum zero found")
	}
}

func RunExample() {
	nums := []int{-4, 17, 49, 4}
	findSubsets(nums)
}
