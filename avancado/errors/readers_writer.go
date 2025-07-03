package erros

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type MyWriter struct {
	w io.Writer
}

func (mw MyWriter) Write(b []byte) (int, error) {
	for i, bb := range b {
		b[i] = bb + 1
	}
	return mw.w.Write(b)
}

func ReadersWriter() {
	str := "hello world"
	reader := strings.NewReader(str)
	buffer := make([]byte, 2)
	n, err := reader.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bytes read:", n)
	fmt.Println("Buffer contents:", string(buffer[:n]))
}

func ReadersWriterLoop() {
	str := "hello world"
	reader := strings.NewReader(str)
	writer := MyWriter{w: os.Stdout}
	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		_, _ = writer.Write(buffer[:n])

		fmt.Println("Bytes read:", n)
		fmt.Println("Buffer contents:", string(buffer[:n]))
	}
}
