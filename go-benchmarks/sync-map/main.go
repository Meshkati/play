package main

import (
	"sync"
)

func SyncMap(n int) {
	var m sync.Map
	for i := 0; i < n; i++ {
		_, _ = m.LoadOrStore(i, i)
	}
}

func RegularMap(n int) {
	mu := sync.RWMutex{}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		mu.Lock()
		m[i] = i
		mu.Unlock()
	}
}

func SyncMapReadWrite(k, n int) {
	var m sync.Map
	wg := sync.WaitGroup{}
	wg.Add(1 + k)
	go func() {
		for i := 0; i < n; i++ {
			_, _ = m.LoadOrStore(i, i)
		}
		wg.Done()
	}()
	for i := 0; i < k; i++ {
		go func() {
			for i := 0; i < n; i++ {
				_, _ = m.Load(i)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func RegularMapReadWrite(k, n int) {
	mu := sync.RWMutex{}
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	wg.Add(1 + k)
	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
		wg.Done()
	}()
	for i := 0; i < k; i++ {
		go func() {
			for i := 0; i < n; i++ {
				mu.RLock()
				_, _ = m[i]
				mu.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
}
