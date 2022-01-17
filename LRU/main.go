package main

import (
	"container/list"
)

type LRUCache struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

type pair struct {
	k, v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if el, ok := c.Keys[key]; ok {
		c.List.MoveToFront(el)
		return el.Value.(pair).v
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if el, ok := c.Keys[key]; ok {
		el.Value = pair{k: key, v: value}
		c.List.MoveToFront(el)
	} else {
		el := c.List.PushFront(pair{
			k: key,
			v: value,
		})
		c.Keys[key] = el
	}
	if c.List.Len() > c.Cap {
		el := c.List.Back()
		c.List.Remove(el)
		delete(c.Keys, el.Value.(pair).k)
	}

}
func main() {

}
