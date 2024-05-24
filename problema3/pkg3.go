package problema3

import (
	"fmt"
)

// findSubsetsUtil is the utility function that uses backtracking to find subsets that sum to zero
func findSubsetsUtil(nums []int, index int, currentSubset []int, currentSum int) bool {
	// Base case: if the current sum is zero and the subset is not empty, we found a valid subset
	if currentSum == 0 && len(currentSubset) > 0 {
		fmt.Println("Subset found:", currentSubset)
		return true
	}

	// Recur for all remaining elements that have not been processed yet
	for i := index; i < len(nums); i++ {
		// Include the current element in the subset
		currentSubset = append(currentSubset, nums[i])

		// Check if this inclusion results in a subset sum of zero
		if findSubsetsUtil(nums, i+1, currentSubset, currentSum+nums[i]) {
			return true
		}

		// Backtrack: remove the current element from the subset
		currentSubset = currentSubset[:len(currentSubset)-1]
	}

	return false
}

// findSubsets finds and prints subsets of nums that sum to zero
func findSubsets(nums []int) {
	currentSubset := []int{}
	if !findSubsetsUtil(nums, 0, currentSubset, 0) {
		fmt.Println("No subset with sum zero found")
	}
}

func RunExample() {
	// Example input: {−7, −3, −2, 5, 8}
	nums := []int{-479, 24, -58, 11, 31, -74, 4, 53, 12, 50, -2, -40, -81}

	findSubsets(nums)
}
