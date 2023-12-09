package resp

// Reply is the interface of redis serialization protocol message
// 主要代表一类数据，将回复内容转换成字节
type Reply interface {
	ToBytes() []byte
}
