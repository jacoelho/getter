package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

var (
	typeNames = flag.String("type", "", "comma-separated list of type names; must be set")
)

const (
	usage = `Usage of getter:
getter [flags] -type T [directory]
getter [flags] -type T files...
Options:
`
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	wantTypes := strings.Split(*typeNames, ",")

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	pkgs, err := loadPackage(args...)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	imports := make(map[string]*packages.Package)
	for k, v := range pkgs.Imports {
		imports[k] = v
	}

	for _, t := range wantTypes {
		obj := pkgs.Types.Scope().Lookup(t)
		if obj == nil {
			continue
		}

		v, err := New(obj, imports)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}

		f, err := v.Generate()
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}

		goFile := os.Getenv("GOFILE")
		ext := filepath.Ext(goFile)
		baseFilename := goFile[0 : len(goFile)-len(ext)]
		targetFilename := baseFilename + "_" + strings.ToLower(t) + "_gen.go"

		if err := f.Save(targetFilename); err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}

func loadPackage(paths ...string) (*packages.Package, error) {
	cfg := &packages.Config{
		Mode:  packages.NeedTypes | packages.NeedImports | packages.NeedFiles | packages.NeedName,
		Tests: false,
		Env:   os.Environ(),
	}
	pkgs, err := packages.Load(cfg, paths...)
	if err != nil {
		return nil, fmt.Errorf("loading packages for inspection: %w", err)
	}

	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	return pkgs[0], nil
}
