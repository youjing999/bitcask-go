package index

import (
	"bitcask-go/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBTree_Put(t *testing.T) {
	bt := NewBTree()
	res := bt.Put(nil, &data.LogRecords{Fid: 1, Offset: 100})

	assert.True(t, res)

	res2 := bt.Put([]byte("aa"), &data.LogRecords{Fid: 1, Offset: 2})
	assert.True(t, res2)
}

func TestBTree_Get(t *testing.T) {
	bt := NewBTree()
	res := bt.Put(nil, &data.LogRecords{Fid: 1, Offset: 100})
	assert.True(t, res)
	get := bt.Get(nil)
	assert.Equal(t, uint32(1), get.Fid)
	assert.Equal(t, int64(100), get.Offset)

	res2 := bt.Put([]byte("aa"), &data.LogRecords{Fid: 1, Offset: 2})
	assert.True(t, res2)
	get2 := bt.Get([]byte("aa"))

	assert.Equal(t, uint32(1), get2.Fid)
	assert.Equal(t, int64(2), get2.Offset)

	t.Log(get2)

}

func TestBTree_Delete(t *testing.T) {
	bt := NewBTree()
	res := bt.Put(nil, &data.LogRecords{Fid: 1, Offset: 100})
	assert.True(t, res)
	d1 := bt.Delete(nil)
	assert.True(t, d1)

	res2 := bt.Put([]byte("aa"), &data.LogRecords{Fid: 1, Offset: 2})
	assert.True(t, res2)
	d2 := bt.Delete([]byte("aa"))
	assert.True(t, d2)
}
