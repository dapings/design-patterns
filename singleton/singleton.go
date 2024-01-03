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

// 饿汉式

var Instance = new(singleton)

// 懒汉式
// 方式一：sync.Once

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

// 方式二：sync.Mutex

var (
	instancesMutex *singleton
	mutex          sync.Mutex
)

func GetInstanceMutex() Singleton {
	// 同步化(锁保护)实际上只在instance变量，第一次被赋值之前才有用。
	// 每次在调用获取实例时，都需要执行锁。
	mutex.Lock()
	if instancesMutex == nil {
		instancesMutex = new(singleton)
	}
	mutex.Unlock()

	return instancesMutex
}

func GetInstanceMutexDoubleCheck() Singleton {
	// 同步化(锁保护)实际上只在instance变量，第一次被赋值之前才有用。
	// 每次在调用获取实例时，都需要执行锁。
	// mutex.Lock()
	// if instancesMutex == nil {
	// 	instancesMutex = new(singleton)
	// }
	// mutex.Unlock()

	// NOTE: 在instance变量有值后，同步化变成一个不必要的瓶颈，使用双重检查，去掉这个小小的额外开销！
	// 双重检查中，在实例化后，锁永远不会被执行。
	// 第一次检查
	if instancesMutex == nil {
		// 可能有多个G同时到达，避免多个G同时实例化singleton。
		mutex.Lock()
		// 每个时刻只会有一个G
		// 第二次检查
		if instancesMutex == nil {
			instancesMutex = new(singleton)
		}
		mutex.Unlock()
	}

	return instancesMutex
}
