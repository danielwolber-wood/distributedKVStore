package distributedKVStore

import "os"

type TransactionLogger interface {
	WriteDelete(key string)
	WritePut(key, value string)
}

type EventType int

const (
	EventDelete = iota
	EventPut
)

type FileTransaction struct {
	Id        uint64
	EventType EventType
	Key       string
	Value     string
}

type FileTransactionLogger struct {
	events chan<- FileTransaction
	errors <-chan error
	file   *os.File
	lastId uint64
}

func (f *FileTransactionLogger) WritePut(key, value string) {
	f.events <- FileTransaction{Id: f.lastId + 1, EventType: EventPut, Key: key, Value: value}
}

func (f *FileTransactionLogger) WriteDelete(key string) {
	f.events <- FileTransaction{Id: f.lastId + 1, EventType: EventDelete, Key: key, Value: ""}
}

func (f *FileTransactionLogger) Err() <-chan error {
	return f.errors
}
