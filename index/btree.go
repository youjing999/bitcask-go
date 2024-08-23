package index

import (
	"bitcask-go/data"
	"github.com/google/btree"
	"sync"
)

// BTree 索引，主要封装了btree kv
// https://githun.com/google/btree
// 这个数据结构在多个goroutine中write操作不安全，read安全
type BTree struct {
	tree *btree.BTree

	lock *sync.RWMutex
}

// NewBTree 初始化BTree 索引结构
func NewBTree() *BTree {
	return &BTree{
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

func (bt *BTree) Put(key []byte, pos *data.LogRecords) bool {
	it := &Item{key: key, pos: pos}

	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(it)
	bt.lock.Unlock()
	return true

}
func (bt *BTree) Get(key []byte) *data.LogRecords {
	it := &Item{key: key}
	getBTreeItem := bt.tree.Get(it)
	if getBTreeItem == nil {
		return nil
	}

	return getBTreeItem.(*Item).pos

}
func (bt *BTree) Delete(key []byte) bool {
	it := &Item{key: key}
	bt.lock.Lock()
	oldItem := bt.tree.Delete(it)
	bt.lock.Unlock()
	if oldItem == nil {
		return false
	}
	return true

}
