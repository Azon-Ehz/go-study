package main

import "fmt"

type MyWriter interface {
	Writer(string) error
}
type MyCloser interface {
	Close() error
}

type writeCloser struct {
	MyWriter
}

type fileWriter struct {
	filePath string
}

type dataBaseWriter struct {
	host string
	port string
	db   string
}

func (fw *fileWriter) Writer(string) error {
	fmt.Println("file writer created")
	return nil
}
func (fw *writeCloser) Close() error {
	fmt.Println("file writer closed")
	return nil
}

func (dw *dataBaseWriter) Writer(string) error {
	fmt.Println("file writer dataBase")
	return nil
}

func main() {
	var mw MyWriter = &writeCloser{
		&dataBaseWriter{},
	}
	mw.Writer("hello world")
}
