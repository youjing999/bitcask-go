package data

// LogRecords 数据内存索引，主要是描述数据在磁盘上的位置
type LogRecords struct {
	Fid    uint32
	Offset int64
}
