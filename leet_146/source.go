package leet_146

import (
	"container/list"
)

type LRUCache struct {
	table *hashTable
}

func Constructor(capacity int) LRUCache {
	return LRUCache{table: newHashTable(capacity)}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.table.Get(key)
	if !ok {
		return -1
	}
	return v
}

func (this *LRUCache) Put(key int, value int) {
	this.table.Set(key, value)
}

type hashTable struct {
	capacity int
	length   int
	buf      []*list.List
	lruQ     *list.List
}

const maxCapacity = 50000

func newHashTable(capacity int) *hashTable {
	if capacity > maxCapacity {
		capacity = maxCapacity
	}
	return &hashTable{capacity: capacity, buf: make([]*list.List, capacity), lruQ: list.New()}
}

// key must be positive
func (ht *hashTable) hash(key int) int {
	return key % ht.capacity
}

func (ht *hashTable) Full() bool {
	return ht.lruQ.Len() >= ht.capacity
}

func (ht *hashTable) Get(key int) (int, bool) {
	idx := ht.hash(key)
	l := ht.buf[idx]
	if nil == l {
		return -1, false
	}
	e := ht.findElem(key, l)
	if e == nil {
		return -1, false
	}
	val := ht.findVal(e)
	if nil == val {
		return -1, false
	}
	ht.lruQ.MoveToFront(val.self)
	return val.value, true
}

func (ht *hashTable) findElem(key int, l *list.List) *list.Element {
	if l == nil {
		return nil
	}
	for e := l.Front(); e != nil; e = e.Next() {
		data, ok := e.Value.(*pair)
		if ok && data != nil && data.key == key {
			return e
		}
	}
	return nil
}

func (ht *hashTable) findVal(e *list.Element) *pair {
	if e == nil {
		return nil
	}
	data, ok := e.Value.(*pair)
	if ok && data != nil {
		return data
	}
	return nil
}

func (ht *hashTable) clearCache() {
	if !ht.Full() {
		return
	}
	e := ht.lruQ.Back()
	if e == nil {
		return
	}
	v := ht.lruQ.Remove(e)
	p, ok := v.(*pair)
	if !ok || p == nil {
		return
	}
	idx := ht.hash(p.key)
	if ht.buf[idx] == nil {
		return
	}
	e = ht.findElem(p.key, ht.buf[idx])
	if nil == e {
		return
	}
	ht.buf[idx].Remove(e)
}

func (ht *hashTable) Set(key, value int) {
	idx := ht.hash(key)
	if ht.buf[idx] == nil {
		ht.buf[idx] = list.New()
	}
	e := ht.findElem(key, ht.buf[idx])
	if e != nil {
		p, _ := e.Value.(*pair)
		p.value = value
		ht.lruQ.MoveToFront(p.self)
		return
	}

	ht.clearCache()

	p := &pair{
		key:   key,
		value: value,
	}
	ht.buf[idx].PushBack(p)
	p.self = ht.lruQ.PushFront(p)
}

type pair struct {
	key   int
	value int
	self  *list.Element
}
