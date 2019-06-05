package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"os"
)

// Reader is title reader of wikipedia's title in ns0.
type Reader struct {
	f  *os.File
	br *bufio.Reader
}

// NewReader creates a new reader.
func NewReader(name string) (*Reader, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	zr, err := gzip.NewReader(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	br := bufio.NewReader(zr)
	return &Reader{
		f:  f,
		br: br,
	}, nil
}

// Close closes reader.
func (r *Reader) Close() error {
	var err error
	if r.f != nil {
		err = r.f.Close()
	}
	return err
}

func (r *Reader) readLine() (string, error) {
	b, _, err := r.br.ReadLine()
	if err != nil {
		if err == io.EOF && len(b) > 0 {
			return string(b), nil
		}
		return "", err
	}
	return string(b), nil
}

// ReadTitle read a title.
func (r *Reader) ReadTitle() (string, error) {
	for {
		s, err := r.readLine()
		if err != nil {
			return "", err
		}
		if s == "" {
			continue
		}
		// FIXME: decode or filter titles.
		return s, nil
	}
}
