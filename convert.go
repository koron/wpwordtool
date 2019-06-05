package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/koron-go/janorm"
	"github.com/koron-go/trietree"
)

type Converter struct {
	dt *trietree.DTree
	ws []string
}

type stringFilter func(string) bool

func NewConverter() *Converter {
	return &Converter{
		dt: &trietree.DTree{},
	}
}

func (c *Converter) Load(name string, filter stringFilter) error {
	r, err := NewReader(name)
	if err != nil {
		return err
	}
	defer r.Close()

	for {
		s, err := r.ReadTitle()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		s = janorm.Normalize(s)
		if s == "" {
			continue
		}
		if filter != nil && !filter(s) {
			continue
		}
		id := c.dt.Put(s)
		if id <= 0 {
			continue
		}
		id--
		if id < len(c.ws) {
			c.ws[id] = s
		} else {
			c.ws = append(c.ws, s)
		}
	}
	return nil
}

func (c *Converter) saveTree(name string) error {
	st := trietree.Freeze(c.dt)
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	err = st.Write(f)
	if err != nil {
		return err
	}
	return nil
}

func (c *Converter) saveWords(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, s := range c.ws {
		fmt.Fprintln(f, s)
	}
	return nil
}

func (c *Converter) Save(basename string) error {
	err := c.saveTree(basename + ".stt")
	if err != nil {
		return err
	}
	err = c.saveWords(basename + ".stw")
	if err != nil {
		os.Remove(basename + ".stt")
		return err
	}
	return nil
}

func jaFilter(s string) bool {
	rs := []rune(s)
	return len(rs) >= 2
}

var rxEn = regexp.MustCompile(`^[A-Za-z][A-Za-z]+$`)

func enFilter(s string) bool {
	return rxEn.MatchString(s)
}

// Convert converts wikipedia's word files to static trietree file.
func Convert(inJa, inEn string, out string) error {
	c := NewConverter()

	log.Printf("loading jawiki: %s", inJa)
	err := c.Load(inJa, jaFilter)
	if err != nil {
		return err
	}

	log.Printf("loading enwiki: %s", inEn)
	err = c.Load(inEn, enFilter)
	if err != nil {
		return err
	}

	log.Printf("writing")
	err = c.Save(out)
	if err != nil {
		return err
	}
	log.Printf("done")
	return nil
}

func runConvert(fs *flag.FlagSet, args []string) error {
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	err = Convert(
		"tmp/jawiki-20190520-all-titles-in-ns0.gz",
		"tmp/enwiki-20190520-all-titles-in-ns0.gz",
		defaultSTreeFile)
	if err != nil {
		return err
	}
	return nil
}
