package singleton

import "sync"

// Singleton -> holds the desired single instance
type singleton struct {
	Data int
}

var (
	once     sync.Once
	instance *singleton
)

// GetInstance - returns a pointer to singleton
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
