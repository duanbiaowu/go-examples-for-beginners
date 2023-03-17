package idiom

func Generate(start, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			ch <- i
		}
	}(ch)

	return ch
}
