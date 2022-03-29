package main

// DLinkNode 双向链表节点
type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

type LRUCache struct {
	size       int
	cap        int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

// InitDLinkNode 初始化双向链表节点
func InitDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:  0,
		cap:   capacity,
		cache: make(map[int]*DLinkNode),
		// 头尾两个哨兵节点置为0
		head: InitDLinkNode(0, 0),
		tail: InitDLinkNode(0, 0),
	}
	// 一开始只有这两个节点
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	// 先判断hash表中有没有这个元素
	if _, ok := this.cache[key]; !ok {
		return -1
	} else {
		// 如果有就将它移动到链表头
		node := this.cache[key]
		this.moveToHead(node)
		// 并返回它的key
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 如果hash表中没有该元素
	if _, ok := this.cache[key]; !ok {
		// 用传进来的key和value初始化一个节点
		node := InitDLinkNode(key, value)
		// 添加到hash表
		this.cache[key] = node
		// 添加到链表头
		this.addToHead(node)
		this.size++
		// 查看缓存是否到达上限
		if this.size > this.cap {
			// 缓存到达上限就删掉尾结点
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		// 如果元素已经存在就更新值
		node := this.cache[key]
		node.value = value
		// 移动到头结点的位置
		this.moveToHead(node)
	}
}

// 双向链表的相关操作
func (this *LRUCache) addToHead(node *DLinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}
func (this *LRUCache) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) moveToHead(node *DLinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}
func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
