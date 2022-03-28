package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("hellu~ this text will be split in eight character chunks"))
	wc.Close()

	// convert WriterCloser interface to BufferedWriterCloser struct
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// convert interface to wrong type
	wrongConversion, ok := wc.(io.Reader)
	if ok {
		fmt.Println(wrongConversion)
	} else {
		fmt.Println("Conversion failed")
	}

	emptyInterface()
}

func emptyInterface() {
	var myObject interface{} = NewBufferedWriterCloser()
	if wc, ok := myObject.(WriterCloser); ok {
		wc.Write([]byte("hellu~ this text will be split in eight character chunks"))
		wc.Close()
	}

	r, ok := myObject.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

// no clue what happening bellow here
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
