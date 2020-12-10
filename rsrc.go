package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/akavel/rsrc/rsrc"
)

var usage = `USAGE:

%s [-manifest FILE.exe.manifest] [-ico FILE.ico[,FILE2.ico...]] -o FILE.syso
  Generates a .syso file with specified resources embedded in .rsrc section,
  aimed for consumption by Go linker when building Win32 excecutables.

The generated *.syso files should get automatically recognized by 'go build'
command and linked into an executable/library, as long as there are any *.go
files in the same directory.

OPTIONS:
`

func main() {
	//TODO: allow in options advanced specification of multiple resources, as a tree (json?)
	//FIXME: verify that data file size doesn't exceed uint32 max value
	var fnamein, fnameico, fnameout, arch string
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.StringVar(&fnamein, "manifest", "", "path to a Windows manifest file to embed")
	flags.StringVar(&fnameico, "ico", "", "comma-separated list of paths to .ico files to embed")
	flags.StringVar(&fnameout, "o", "rsrc.syso", "name of output COFF (.res or .syso) file")
	flags.StringVar(&arch, "arch", "386", "architecture of output file - one of: 386, amd64, [EXPERIMENTAL: arm, arm64]")
	_ = flags.Parse(os.Args[1:])
	if fnameout == "" || (fnamein == "" && fnameico == "") {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flags.PrintDefaults()
		os.Exit(1)
	}

	err := rsrc.Embed(fnameout, arch, fnamein, fnameico)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
