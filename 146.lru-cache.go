package main

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
type LRUCache struct {
	valMap   map[int]int
	lruList  DoublyLinkedList[int]
	capacity int
}

func ___Constructor(capacity int) LRUCache {
	return LRUCache{
		valMap:   make(map[int]int),
		lruList:  DoublyLinkedList[int]{},
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.valMap[key]; ok {
		this.lruList.RemoveVal(key)
		this.lruList.Prepend(key)
		return val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.valMap[key]; ok {
		this.lruList.RemoveVal(key)
		this.lruList.Prepend(key)
	} else {
		if this.lruList.length == this.capacity {
			dk := this.lruList.Pop()
			delete(this.valMap, dk)
		}
		this.lruList.Prepend(key)
	}
	this.valMap[key] = value
}
