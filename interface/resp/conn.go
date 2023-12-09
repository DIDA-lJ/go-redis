package resp

type Connection interface {
	Write([]byte) error
	// used for multi database
	GetDBIndex() int
	SelectDB(int)
}
