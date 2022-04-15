package main

import (
	"math/rand"
	"sync"
	"testing"
)

type SetMutex struct {
	set map[float64]struct{}
	mu  *sync.Mutex
	mu2 *sync.RWMutex
}

func NewSet() *SetMutex {
	return &SetMutex{
		set: map[float64]struct{}{},
		mu:  &sync.Mutex{},
		mu2: &sync.RWMutex{},
	}
}

// пишем в карту, причем key = действительные числа, а value = пустые структуры
func (s *SetMutex) Get(x float64) {
	s.mu.Lock()
	s.set[x] = struct{}{}
	s.mu.Unlock()
}

// читаем из карты(структуры)
func (s *SetMutex) Set(x float64) {
	s.mu.Lock()
	_ = s.set[x]
	s.mu.Unlock()
}

func (s *SetMutex) Get2(x float64) {
	s.mu2.Lock()
	s.set[x] = struct{}{}
	s.mu2.Unlock()
}

func (s *SetMutex) Set2(x float64) {
	s.mu2.RLock()
	_ = s.set[x]
	s.mu2.RUnlock()
}

// sync.Mutex для разных вариантов использования: 10% запись, 90% чтение
func BenchmarkMutex10w_90r(b *testing.B) {
	var wg sync.WaitGroup
	structSet := NewSet()
	// время на создание structSet не учитываем
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(j int) {
				structSet.Get(rand.Float64())
				wg.Done()
			}(i)
		}
		for k := 0; k < 90; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.Set(rand.Float64())
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}

// sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение
func BenchmarkRWMutex10w_90r(b *testing.B) {
	var wg sync.WaitGroup
	structSet := NewSet()
	// время на создание structSet не учитываем
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(j int) {
				structSet.Get2(rand.Float64())
				wg.Done()
			}(i)
		}
		for k := 0; k < 90; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.Set2(rand.Float64())
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}

// sync.Mutex для разных вариантов использования: 50% запись, 50% чтение
func BenchmarkMutex50w_50r(b *testing.B) {
	var wg sync.WaitGroup
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 50; j++ {
			wg.Add(1)
			go func(j int) {
				structSet.Get(rand.Float64())
				wg.Done()
			}(i)
		}
		for k := 0; k < 50; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.Set(rand.Float64())
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}

// sync.RWMutex для разных вариантов использования: 50% запись, 50% чтение
func BenchmarkRWMutex50w_50r(b *testing.B) {
	var wg sync.WaitGroup
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 50; j++ {
			wg.Add(1)
			go func(j int) {
				structSet.Get2(rand.Float64())
				wg.Done()
			}(i)
		}
		for k := 0; k < 50; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.Set2(rand.Float64())
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}

// sync.Mutex для разных вариантов использования: 90% запись, 10% чтение
func BenchmarkMutex90w_10r(b *testing.B) {
	var wg sync.WaitGroup
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 90; j++ {
			wg.Add(1)
			go func(j int) {
				structSet.Get(rand.Float64())
				wg.Done()
			}(i)
		}
		for k := 0; k < 10; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.Set(rand.Float64())
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}

// sync.RWMutex для разных вариантов использования: 90% запись, 10% чтение
func BenchmarkRWMutex90w_10r(b *testing.B) {
	var wg sync.WaitGroup
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 90; j++ {
			wg.Add(1)
			go func(j int) {
				structSet.Get2(rand.Float64())
				wg.Done()
			}(i)
		}
		for k := 0; k < 10; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.Set2(rand.Float64())
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}
