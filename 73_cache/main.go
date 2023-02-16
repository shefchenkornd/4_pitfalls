package main

import (
	"fmt"
	"sync"
	"time"
)

// Задача: написать простенькую реализацию cache используя map[string]string

type Cacher interface {
	Get(k string) string
	Set(k string, v string)
	Delete(k string)
}

type Cache struct {
	mx sync.Mutex
	m  map[string]string
	t ttl
}

func (c *Cache) Get(k string) string {
	v, ok := c.m[k]
	if !ok {
		return ""
	}

	return v
}
type ttl int32
func (c *Cache) Set(k string, v string) {
	c.mx.Lock()
	c.m[k] = v
	c.mx.Unlock()
}

func (c *Cache) Delete(k string) {
	_, ok := c.m[k]
	if ok {
		c.mx.Lock()
		delete(c.m, k)
		c.mx.Unlock()
	}
}

func New() Cacher {
	return &Cache{
		mx: sync.Mutex{},
		m:  make(map[string]string),
	}
}

func main() {
	// var cache Cacher = &Cache{} // panic: assignment to entry in nil map
	cache := New() // вот так правильно!

	

	time.After(1*time.Second)
	time.Now().Unix()

	cache.Set("a", "b")
	fmt.Println("cache.Get(a):", cache.Get("a"))

	fmt.Println("cache.Delete(a)")
	cache.Delete("a")
}
