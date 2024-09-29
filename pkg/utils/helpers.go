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

func Fill(text, color string) string {
	return color + text + "\033[0m"
}
