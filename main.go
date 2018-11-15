package main

import (
	"fmt"
	"github.com/bigkucha/Algorithm/cache"
)

func main() {
	lru := cache.NewLRUCache(10)
	lru.Set("name", "Joe")
	name, exist, err := lru.Get("name")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !exist {
		fmt.Println("不存在")
		return
	}
	fmt.Println("缓存存在，值为:", name)
}
