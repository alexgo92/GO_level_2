package main

import (
	"sync"
	"testing"
)

/*sync.Mutex для разных вариантов использования: 10% запись, 90% чтение;
50% запись, 50% чтение; 90% запись, 10% чтениее*/
type SetMutex struct {
	set map[float64]struct{}
	sync.Mutex
}

// в этуу структуру будем осуществлять чтение
var Some struct{}

func NewSet() *SetMutex {
	return &SetMutex{
		set: map[float64]struct{}{},
	}
}

// пишем в карту, причем key = действительные числа, а value = пустые структуры
func (s *SetMutex) Add(x float64) {
	s.Lock()
	s.set[x] = struct{}{}
	s.Unlock()
}

// читаем в Some из карты(структуры)
func (s *SetMutex) ReadSet(x float64) {
	s.Lock()
	Some = s.set[x]
	s.Unlock()
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
				structSet.Add(float64(j))
				wg.Done()
			}(i)
		}
		for k := 0; k < 90; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.ReadSet(float64(k))
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
				structSet.Add(float64(j))
				wg.Done()
			}(i)
		}
		for k := 0; k < 50; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.ReadSet(float64(k))
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
				structSet.Add(float64(j))
				wg.Done()
			}(i)
		}
		for k := 0; k < 10; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.ReadSet(float64(k))
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}

/*sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение;
50% запись, 50% чтение; 90% запись, 10% чтениее*/
type SetRWMutex struct {
	set map[float64]struct{}
	sync.RWMutex
}

func NewRWSet() *SetRWMutex {
	return &SetRWMutex{
		set: map[float64]struct{}{},
	}
}

// при записи используем Lock() и Unlock()
func (s *SetRWMutex) AddRW(x float64) {
	s.Lock()
	s.set[x] = struct{}{}
	s.Unlock()
}

// при чтении используем RLock() и RUnlock()
func (s *SetRWMutex) ReadSetRW(x float64) {
	s.RLock()
	Some = s.set[x]
	s.RUnlock()
}

// sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение
func BenchmarkRWMutex10w_90r(b *testing.B) {
	var wg sync.WaitGroup
	structSet := NewRWSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(j int) {
				structSet.AddRW(float64(j))
				wg.Done()
			}(i)
		}
		for k := 0; k < 90; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.ReadSetRW(float64(k))
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}

//sync.RWMutex для разных вариантов использования: 50% запись, 50% чтение
func BenchmarkRWMutex50w_50r(b *testing.B) {
	var wg sync.WaitGroup
	structSet := NewRWSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		for j := 0; j < 50; j++ {
			wg.Add(1)
			go func(j int) {
				structSet.AddRW(float64(j))
				wg.Done()
			}(i)
		}
		for k := 0; k < 50; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.ReadSetRW(float64(k))
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}

// sync.RWMutex для разных вариантов использования: 90% запись, 10% чтение
func BenchmarkRWMutex90w_10r(b *testing.B) {
	var wg sync.WaitGroup
	structSet := NewRWSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 90; j++ {
			wg.Add(1)
			go func(j int) {
				structSet.AddRW(float64(j))
				wg.Done()
			}(i)
		}
		for k := 0; k < 10; k++ {
			wg.Add(1)
			go func(k int) {
				structSet.ReadSetRW(float64(k))
				wg.Done()
			}(k)
		}
	}
	wg.Wait()
}
