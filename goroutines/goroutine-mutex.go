package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	m  map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	s.m[key]++
	s.mu.Unlock()
}

// demo using waitGroup to wait all goroutines to finish
// use wg.Add() wg.Done() wg.Wait()
func (s *SafeCounter) Dec(wg *sync.WaitGroup, key string) {
	s.mu.Lock()
	s.m[key]--
	s.mu.Unlock()
	wg.Done()
}

func (s *SafeCounter) Value(key string) int {
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
	}()
	return s.m[key]
}

func main() {
	key := "hello"
	s := &SafeCounter{m: make(map[string]int)}
	for i := 0; i < 20; i++ {
		go s.Inc(key)
	}
	time.Sleep(time.Second)

	var wg sync.WaitGroup
	fmt.Println(s.Value(key))
	fmt.Println("----------------")
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			s.Dec(&wg, key)
		}()
	}
	wg.Wait()
	fmt.Println(s.Value(key))
}
