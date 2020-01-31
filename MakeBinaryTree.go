package main

func MakeBT(nums []int) []int {
	if nums == nil {
		return nil
	}
	slow := 0
	fast := 1

	for ; slow < len(nums)/2; slow++ {
		//sum := 0
		for fast = slow*2 + 1; fast < len(nums); fast++ {

		}
	}
	return nil
}
