// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// This program generates bits_tables.go.

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
)

var header = []byte(`// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by go run make_tables.go. DO NOT EDIT.

package bits

`)

func main() {
	buf := bytes.NewBuffer(header)

	gen(buf, "rev8tab", rev8)
	// add more tables as needed

	out, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("bits_tables.go", out, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func gen(w io.Writer, name string, f func(uint8) uint8) {
	fmt.Fprintf(w, "var %s = [256]uint8{", name)
	for i := 0; i < 256; i++ {
		if i%16 == 0 {
			fmt.Fprint(w, "\n\t")
		} else {
			fmt.Fprint(w, " ")
		}
		fmt.Fprintf(w, "%#02x,", f(uint8(i)))
	}
	fmt.Fprint(w, "\n}\n\n")
}

func rev8(x uint8) (r uint8) {
	for i := 8; i > 0; i-- {
		r = r<<1 | x&1
		x >>= 1
	}
	return
}
