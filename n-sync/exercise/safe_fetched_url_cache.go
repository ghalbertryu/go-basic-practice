package main

import "sync"

type SafeFetchedUrlCache struct {
	mu   sync.Mutex
	urls map[string]int
}

func (cache *SafeFetchedUrlCache) add(url string) {
	cache.mu.Lock()
	cache.urls[url]++
	cache.mu.Unlock()
}

func (cache *SafeFetchedUrlCache) has(url string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	return cache.urls[url] > 0
}
