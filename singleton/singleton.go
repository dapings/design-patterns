package singleton

import (
	"sync"
)

// Singleton 单例模式接口，通过此接口可以避免 GetInstance 返回一个包私有类型的指针
type Singleton interface {
	bar()
}

type singleton struct {
	data string
}

func (s singleton) bar() {}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{data: "Hello, Singleton!"}
	})

	return instance
}
