package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/koron-go/janorm"
	"github.com/koron-go/trietree"
)

type Matcher struct {
	st *trietree.STree
	ws []string
}

func LoadMatcher(basename string) (*Matcher, error) {
	st, err := loadSTree(basename + ".stt")
	if err != nil {
		return nil, err
	}
	ws, err := loadWords(basename + ".stw")
	if err != nil {
		return nil, err
	}
	return &Matcher{
		st: st,
		ws: ws,
	}, nil
}

func loadSTree(name string) (*trietree.STree, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	st, err := trietree.Read(f)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func loadWords(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	br := bufio.NewReader(f)
	var ws []string
	for {
		b, _, err := br.ReadLine()
		if len(b) == 0 && err != nil {
			if err == io.EOF {
				return ws, nil
			}
			return nil, err
		}
		ws = append(ws, string(b))
	}
}

func (m *Matcher) MatchAll(s string) []string {
	var list []string
	sc := NewScanner(s)
	m.st.Scan(s, sc)
	for _, id := range sc.Finish() {
		if id <= 0 {
			continue
		}
		id--
		if id < len(m.ws) {
			list = append(list, m.ws[id])
		}
	}
	return list
}

func (m *Matcher) MatchFile(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		b, _, err := br.ReadLine()
		if len(b) == 0 && err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		s := string(b)
		fmt.Print(s)
		s = janorm.Normalize(s)
		for _, w := range m.MatchAll(s) {
			fmt.Print("\t", w)
		}
		fmt.Print("\n")
	}
}

func runMatch(fs *flag.FlagSet, args []string) error {
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	m, err := LoadMatcher(defaultSTreeFile)
	if err != nil {
		return err
	}
	for _, fn := range fs.Args() {
		log.Printf("matching against file:%s", fn)
		err := m.MatchFile(fn)
		if err != nil {
			return err
		}
	}
	return nil
}
