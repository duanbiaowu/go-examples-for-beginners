package creational

type Pool chan struct{}

func NewPool(total int) *Pool {
	p := make(Pool, total)
	for i := 0; i < total; i++ {
		p <- struct{}{}
	}
	return &p
}
