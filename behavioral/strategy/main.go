package main

import "fmt"

type FIFO struct{}

func (l *FIFO) Evict(c *Cache) {
	fmt.Println("evicting by FIFO")
}

type LRU struct{}

func (l *LRU) Evict(c *Cache) {
	fmt.Println("evicting by LRU")
}

func main() {
	lru := &LRU{}
	cache := NewCache(lru)
	cache.Add("hello", "world")
	cache.Add("fromis", "9")
	res := cache.Get("fromis")
	fmt.Println(res)

	cache.Add("love", "bomb")

	fifo := &FIFO{}
	cache.SetEvictionAlgorithm(fifo)
	cache.Add("9 way", "ticket")
}
