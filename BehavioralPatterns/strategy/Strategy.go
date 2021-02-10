package main

import (
	"fmt"
)

// 缓存策略接口
type evictionAlgo interface {
	evict(c *cache)
}


// 缓存策略的实现：fifo
type fifo struct {

}

func(l *fifo)evict(c *cache){
	fmt.Println("Evicting by fifo strtegy")
}

// 缓存策略的实现：fifo
type lru struct {

}

func(l *lru)evict(c *cache){
	fmt.Println("Evicting by lru strtegy")
}

// 缓存策略的实现：lfu
type lfu struct {

}

func(l *lfu)evict(c *cache){
	fmt.Println("Evicting by lfu strtegy")
}

// 在 cache 结构体中选择不同的缓存实现
type cache struct {
	storage map[string]string
	evictionAlgo evictionAlgo  // 组合策略接口
	capacity int
	maxCapacity int
}


func(c *cache) setEviction(e evictionAlgo){ // 设置一个缓存清除策略到cache中
	c.evictionAlgo = e
}

func (c *cache)evict(){
	c.evictionAlgo.evict(c)
	c.capacity--
}


func(c *cache) add(key, value string){
	if c.capacity == c.maxCapacity{ // 如果满了就执行清除策略
		c.evict()
	}
	c.capacity++
	c.storage[key]=value
}


func (c *cache) get(key string){
	delete(c.storage, key)
}


func initCache() *cache{
	storage := make(map[string]string)
	return &cache{
		storage:storage,
		capacity:0,
		maxCapacity:2,
	}
}


func strategyExecutor(evict string){
	cache := initCache()
	switch evict {

	case "lfu":
		lfu := &lfu{}
		cache.setEvictionAlgo(lfu)

	case "lru":
		lru := &lru{}
		cache.setEvictionAlgo(lru)

	case "fifo":
		fifo := &fifo{}
		cache.setEvictionAlgo(fifo)
	}
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")
	cache.add("d", "4")
}


func main() {

	strategyExecutor("lru")
	//strategyExecutor("fifo")
	//strategyExecutor("lru")

}



