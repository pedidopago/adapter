package adapter

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/pedidopago/adapter"
)

//go:embed template.go.tpl
var rawtpl string

type GenerateInput struct {
	Package          string
	SpecFileContents []byte
	Output           io.Writer
}

type Package struct {
	Path string
	Name string
}

type GoFile struct {
	PackageName string
	Imports     []Package
	Adapters    []adapter.Spec
}

func Generate(input GenerateInput) error {
	specs := &adapter.Specs{}
	if _, err := toml.Decode(string(input.SpecFileContents), &specs); err != nil {
		return err
	}
	gof := &GoFile{
		PackageName: input.Package,
		Adapters:    make([]adapter.Spec, 0),
	}
	for i := range specs.Items {
		spec := &specs.Items[i]
		if spec.Src == "" || spec.Dst == "" {
			return fmt.Errorf("spec %d: missing src and/or dst", i)
		}
		for i := range spec.Fields {
			field := &spec.Fields[i]
			if field.To == "" {
				field.To = field.From
			}
		}
		if strings.HasPrefix(spec.Dst, "*") {
			spec.DstIsPtr = true
		}
		if spec.SliceName == "" {
			spec.SliceName = spec.Name + "Slice"
		}
		gof.Adapters = append(gof.Adapters, *spec)
	}
	pkgs := make([]Package, 0, len(specs.Packages))
	for _, pkg := range specs.Packages {
		if strings.Contains(pkg, ";") {
			pkgss := strings.SplitN(pkg, ";", 2)
			pkgs = append(pkgs, Package{
				Name: pkgss[1],
				Path: pkgss[0],
			})
		} else {
			pkgs = append(pkgs, Package{
				Name: "",
				Path: pkg,
			})
		}
	}
	gof.Imports = pkgs
	tpl := template.Must(template.New("gofile").Parse(rawtpl))
	return tpl.Execute(input.Output, gof)
}
