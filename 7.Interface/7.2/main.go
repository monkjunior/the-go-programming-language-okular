package main

import (
	"fmt"
	"io"
	"os"
)

type NewWriter struct {
	bytes int64
	writer io.Writer
}

func (w *NewWriter) Write(p []byte) (int, error){
	w.bytes = int64(len(p))
	return 	w.writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64){
	newWriter := NewWriter{0, w}
	return &newWriter, &newWriter.bytes
}

func main(){
	newWriter, bytesCounter := CountingWriter(os.Stdout)

	fmt.Fprintf(newWriter, "Hello\n")
	fmt.Println("Bytes", *bytesCounter)

	fmt.Fprintf(newWriter, "123456789\n")
	fmt.Println("Bytes", *bytesCounter)
}

