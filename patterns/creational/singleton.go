package creational

import "sync"

var (
	once     sync.Once
	instance Singleton
)

// Singleton struct
type Singleton struct {
}

// NewInstance Singleton
func NewInstance() Singleton {
	once.Do(func() {
		instance = Singleton{}
	})
	return instance
}
