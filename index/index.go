package index

import (
	"bitcask-go/data"
	"bytes"
	"github.com/google/btree"
)

// Indexer 抽象索引接口，后续如果想接入其他数据结构，则直接实现这个接口
type Indexer interface {
	// Put 向索引中添加key 对应的数据位置信息
	Put(key []byte, pos *data.LogRecords) bool
	// Get 根据 key 获取对应索引的位置信息
	Get(key []byte) *data.LogRecords
	// Delete 根据 key 删除对应索引的位置信息
	Delete(key []byte) bool
}

// Item 对tree进行排序需要实现Less
/*
	type Item interface {
	// Less tests whether the current item is less than the given argument.
	//
	// This must provide a strict weak ordering.
	// If !a.Less(b) && !b.Less(a), we treat this to mean a == b (i.e. we can only
	// hold one of either a or b in the tree).
	Less(than Item) bool
}
*/
type Item struct {
	key []byte
	pos *data.LogRecords
}

func (ai *Item) Less(bt btree.Item) bool {
	return bytes.Compare(ai.key, bt.(*Item).key) == -1
}
