package main

import "fmt"

type myWriter interface {
	Writer(string)
}

type myReader interface {
	Reader() string
}

type myReadWriter interface {
	myWriter
	myReader
	ReadWriter()
}

type SreadWriter struct {
}

func (s *SreadWriter) Writer(s2 string) {
	//TODO implement me
	fmt.Println("Writer me")
}

func (s *SreadWriter) Reader() string {
	fmt.Println("Reader me")
	return ""
}

func (s *SreadWriter) ReadWriter() {
	//TODO implement me
	fmt.Println("ReadWriter me")
}

func main() {
	var mrw myReadWriter = &SreadWriter{}
	mrw.ReadWriter()
}
