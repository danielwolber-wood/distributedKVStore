package distributedKVStore

import (
	"fmt"
	"os"
)

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

func (f *FileTransactionLogger) Run() {
	events := make(chan FileTransaction, 16)
	f.events = events
	errors := make(chan error, 1)
	f.errors = errors
	go func() {
		for e := range events {
			f.lastId++
			_, err := fmt.Fprintf(f.file, "%d\t%d\t%s\t%s\n",
				f.lastId, e.EventType, e.Key, e.Value)
			if err != nil {
				errors <- err
			}
		}
	}()
}

func NewFileTransactionLogger(filename string) (FileTransactionLogger, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return FileTransactionLogger{}, err
	}
	return FileTransactionLogger{file: file}, nil
}
