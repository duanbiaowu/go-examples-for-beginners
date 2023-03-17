package mapreduce

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray []string
	for _, val := range arr {
		newArray = append(newArray, fn(val))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray []int
	for _, val := range arr {
		newArray = append(newArray, fn(val))
	}
	return newArray
}

func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, val := range arr {
		sum += fn(val)
	}
	return sum
}

func Filter(arr []int, fn func(n int) bool) []int {
	var newArray []int
	for _, val := range arr {
		if fn(val) {
			newArray = append(newArray, val)
		}
	}
	return newArray
}
