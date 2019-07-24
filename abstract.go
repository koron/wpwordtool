package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// runAbstract extracts only abstracts from Wikipedia's XML.
func runAbstract(fs *flag.FlagSet, args []string) error {
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	for _, fn := range fs.Args() {
		log.Printf("extracting from %s", fn)
		err := filterAbstractFile(os.Stdout, fn)
		if err != nil {
			return err
		}
	}
	return nil
}

func filterAbstractFile(out io.Writer, name string) error {
	var r io.Reader
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	r = f
	if strings.HasSuffix(name, ".gz") {
		r, err = gzip.NewReader(r)
		if err != nil {
			return err
		}
	}
	br := bufio.NewReader(r)
	for {
		b, _, err := br.ReadLine()
		if len(b) == 0 && err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if !bytes.HasPrefix(b, []byte("<abstract>")) || !bytes.HasSuffix(b, []byte("</abstract>")) {
			continue
		}
		b = b[10 : len(b)-11]
		out.Write(b)
		fmt.Fprintln(out)
	}
}
