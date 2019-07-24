package main

import (
	"flag"
	"log"

	subcmd "github.com/koron/go-subcmd"
)

const (
	defaultSTreeFile = "tmp/wikiwords"
)

var cmds = subcmd.Subcmds{
	"convert":  subcmd.Main2(runConvert),
	"match":    subcmd.Main2(runMatch),
	"abstract": subcmd.Main2(runAbstract),
}

func main() {
	flag.Parse()

	err := cmds.RunWithName("wpwordtool", flag.Args())
	if err != nil {
		log.Fatal(err)
	}
}
