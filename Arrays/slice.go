package arrays


func sliceSum (nums []int) int {
	sum := 0;
	
	for _, i:= range nums {
		sum += i
	}
	return sum
}