package iteration

const repeatCount = 5

func Repeat(letter string, times int) (repeated string) {
	if times == 0 {
		times = repeatCount
	}
	for i := 0; i < times; i++ {
		repeated += letter
	}
	return
}
