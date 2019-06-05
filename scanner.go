package main

import "github.com/koron-go/trietree"

type Scanner struct {
	rs []rune

	ids   []int
	lasts []int

	ridx int
}

func NewScanner(s string) *Scanner {
	rs := []rune(s)
	lasts := make([]int, len(rs))
	for i := range lasts {
		lasts[i] = i
	}
	return &Scanner{
		rs:    rs,
		ids:   make([]int, len(rs)),
		lasts: lasts,
	}
}

func (sc *Scanner) ScanReport(ev trietree.ScanEvent) {
	ridx := sc.ridx
	sc.ridx++
	if len(ev.Nodes) == 0 {
		return
	}
	for _, n := range ev.Nodes {
		start := ridx - n.Level + 1
		last := ridx
		if sc.ids[start] == 0 || last > sc.lasts[start] {
			sc.ids[start] = n.ID
			sc.lasts[start] = last
		}
	}
}

func (sc *Scanner) Finish() []int {
	var list []int
	for i := 0; i < len(sc.rs); i++ {
		id := sc.ids[i]
		if id != 0 {
			list = append(list, id)
		}
		i = sc.lasts[i]
	}
	return list
}
