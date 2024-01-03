package singleton

import (
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("expect instance equal, but got not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	done := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)

	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			// block, wait done chan closed
			<-done
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	// 唤醒被阻塞的G，多核情况下会被调度到多个P同时运行，实现并行
	close(done)
	wg.Wait()

	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("expect instance equal, but got not equal")
		}
	}
}

func BenchmarkSingleton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInstanceMutex()
	}
}

func BenchmarkSingletonDoubleCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetInstanceMutexDoubleCheck()
	}
}
