package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pedidopago/adapter/internal/adapter"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "package",
			Aliases: []string{"p"},
			EnvVars: []string{"GOPACKAGE"},
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"out", "o"},
		},
		&cli.StringFlag{
			Name:    "spec",
			Aliases: []string{"file", "f"},
			EnvVars: []string{"SPECFILE", "SPEC"},
		},
	}
	app.Action = cmdwrap(cmdrun)
	if err := app.Run(os.Args); err != nil {
		e := err.(cli.ExitCoder)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(e.ExitCode())
	}
}

func cmdwrap(fn func(c *cli.Context) cli.ExitCoder) cli.ActionFunc {
	return func(c *cli.Context) error {
		return fn(c)
	}
}

func cmdrun(c *cli.Context) cli.ExitCoder {
	outfile := c.String("output")
	if outfile == "" {
		if os.Getenv("GOFILE") == "" {
			return cli.NewExitError("output file not specified", 1)
		}
		outfile = strings.TrimSuffix(os.Getenv("GOFILE"), ".go") + "_adapter.gen.go"
	}
	specfile := c.String("spec")
	inputf, err := os.ReadFile(specfile)
	if err != nil {
		return cli.NewExitError(fmt.Errorf("failed to open %s: %w", inputf, err), 1)
	}
	outf, err := os.Create(outfile)
	if err != nil {
		return cli.NewExitError(fmt.Errorf("failed to create %s: %w", outfile, err), 1)
	}
	if err := adapter.Generate(adapter.GenerateInput{
		Package:          c.String("package"),
		SpecFileContents: inputf,
		Output:           outf,
	}); err != nil {
		outf.Close()
		return cli.NewExitError(err, 2)
	}
	outf.Close()
	if err := exec.Command("go", "fmt", outfile).Run(); err != nil {
		return cli.NewExitError(fmt.Errorf("gofmt %s: %w", outfile, err), 3)
	}
	return nil
}
