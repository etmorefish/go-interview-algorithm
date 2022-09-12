package main

/*
最近最少使用， LRU缓存就是使用这种原理实现，简单的说就是缓存一定量的数据，当超过设定的阈值时就把一些过期的数据删掉。常用于页面置换算法。

请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/
func main() {

}

// 定义双向链表
type BLinkNode struct {
	key, val  int
	pre, next *BLinkNode
}
type LRUCache struct {
	size, capacity int
	cache          map[int]*BLinkNode
	head, tail     *BLinkNode
}

// 初始化一个双向链表
func initBLinkNode(key, val int) *BLinkNode {
	return &BLinkNode{
		key: key,
		val: val,
	}
}

// LRUCache 初始化
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		size:     0,
		cache:    map[int]*BLinkNode{},
		head:     initBLinkNode(0, 0),
		tail:     initBLinkNode(0, 0),
	}
	l.head.next = l.tail // 双链表头尾相连
	l.tail.next = l.head
	return l
}

// 移除节点
func removeNode(node *BLinkNode) {

	node.pre.next = node.next // 上一个节点的下一个指向 指向 当前下一个节点
	node.next.pre = node.pre  // 下一个节点的上一个指向 指向 当前上一个节点
}

// 移除尾部节点
func (this *LRUCache) removeTail() *BLinkNode {
	node := this.tail.pre //获取当前节点（尾部不存数据）
	removeNode(node)
	return node

}

// 添加头节点
func (this *LRUCache) addHead(node *BLinkNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node //原来节点的上一个指向 指向 node
	this.head.next = node     // 改变空头节点的下一个指向 指向 node
}

// 使用的时候 移动到头部
func (this *LRUCache) moveToHead(node *BLinkNode) {
	removeNode(node)
	this.addHead(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok { // 判断key是否在缓存中
		node.val = value      // 是，更新值
		this.moveToHead(node) // 移动到队首
	} else {
		newNode := initBLinkNode(key, value) // 不在缓存中，则新建node
		this.addHead(newNode)                //新节点添加到头部
		this.cache[key] = newNode            // 添加缓存
		this.size++
		if this.size > this.capacity { // 超过容量大小
			tail := this.removeTail()
			this.size--
			delete(this.cache, tail.key) // 删除缓存
		}
	}
}
