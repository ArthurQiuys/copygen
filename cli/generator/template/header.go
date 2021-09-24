package template

import (
	"github.com/switchupcb/copygen/cli/models"
	"github.com/switchupcb/copygen/cli/pkg/interpreter"
)

// Header determines the func to generate header code.
func Header(g *models.Generator) (string, error) {
	if g.Template.Headpath == "" {
		return defaultHeader(g), nil
	} else {
		return interpretHeader(g)
	}
}

// defaultHeader creates the header of the generated file using the default method.
func defaultHeader(g *models.Generator) string {
	var header string

	// package
	header += "// Code generated by github.com/switchupcb/copygen\n"
	header += "// DO NOT EDIT.\n"
	header += "package " + g.Package + "\n"

	// imports
	header += "import (\n"
	for _, iprt := range g.Imports {
		header += "\"" + iprt + "\"\n"
	}
	header += ")"
	return header
}

// interpretHeader creates the header of the generated file using an interpreted template file.
func interpretHeader(g *models.Generator) (string, error) {
	fn, err := interpreter.InterpretFunc(g.Loadpath, g.Template.Funcpath, "generator.Header")
	if err != nil {
		return "", err
	}

	// run the interpreted function.
	return fn(g), nil
}