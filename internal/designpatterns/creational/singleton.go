package creational

import "sync"

var (
	singleton *Singleton
	once      sync.Once
)

type Singleton struct{}

// use sync.Once to initialize the instance once and return the global variable access address.
func GetInstance() *Singleton {

	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
