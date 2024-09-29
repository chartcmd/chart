package utils

func GetClosestNumDivBy(num, threshold int) int {
	prev := num
	for {
		cur := prev + num
		if cur > threshold {
			return prev
		}
		prev = cur
	}
}
