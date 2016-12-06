package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	// "log"
	// "os"

	"compress/gzip"
)

func main() {
	fmt.Println("Hello")

	/*
	 * r, err := gzip.NewReader(os.Stdin)
	 * if err != nil {
	 *     panic(err)
	 * }
	 */

	bs := bytes.NewBufferString("")
	gw := gzip.NewWriter(bs)
	io.WriteString(gw, "hello world")
	gw.Close()

	// fmt.Println("gw: ", gw)

	// io.Copy(os.Stdout, bs)

	// fmt.Println(bs.String())

	gr, err := gzip.NewReader(bs)
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(os.Stdout, gr)
	fmt.Println("xxx")
	fmt.Println("n: ", n, err)

	// bs := make([]byte, 8)

	/*
	 * strings.
	 * w.Write(bs)
	 */

	/*
	 * b := make([]byte, 8)
	 * r.Read(b)
	 * fmt.Println("b: ", b)
	 */

}
