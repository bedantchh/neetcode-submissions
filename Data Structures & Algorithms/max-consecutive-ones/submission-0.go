func findMaxConsecutiveOnes(nums []int) int {
	n:= len(nums)
	maxC:=0
	for i:=0;i<n;i++ {
		count := 0
		for j:=i;j<n;j++ {
			if nums[j] == 0{
				break
			}
			count ++
		}
		if count > maxC {
			maxC = count
		}
	}
	return maxC
}
