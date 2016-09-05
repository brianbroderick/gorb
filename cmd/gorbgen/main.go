package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/lsegal/gorb/codegen"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: gorbgen [options] <package> [package ...]\n\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	boolPtr := flag.Bool("build", false, "builds the extension after generating")
	rootPtr := flag.String("root", "", "root Ruby module to place code under")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	for i := 0; i < flag.NArg(); i++ {
		p := path.Clean(flag.Arg(i))
		gen := codegen.Generator{Path: p, Root: *rootPtr, Build: *boolPtr}
		gen.Generate()
	}
}
