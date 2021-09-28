package interpreter

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/traefik/yaegi/interp"
)

// interpretFunc loads a template package.function into an interpreter.
func interpretFunc(loadpath string, templatepath string, symbol string) (func(interface{}) string, error) {
	i, err := interpretFile(loadpath, templatepath, symbol)
	if err != nil {
		return nil, err
	}

	v, err := i.Eval(symbol)
	if err != nil {
		return nil, fmt.Errorf("An error occured loading a template function.\n%v", err)
	}
	return v.Interface().(func(interface{}) string), nil
}

// interpretFile creates an interpreter loaded with the specified file.
func interpretFile(loadpath string, templatepath, symbol string) (*interp.Interpreter, error) {
	// determine actual filepath
	absfilepath, err := filepath.Abs(loadpath)
	if err != nil {
		return nil, err
	}
	absfilepath = path.Join(filepath.Dir(absfilepath), templatepath)

	// read the file
	file, err := os.ReadFile(absfilepath)
	if err != nil {
		return nil, fmt.Errorf("The specified template file for the template function %v doesn't exist: %v\n", symbol, absfilepath)
	}
	source := string(file)

	// setup the interpreter
	goCache, err := os.UserCacheDir()
	if err != nil {
		return nil, fmt.Errorf("An error occurred loading the template file. Is the GOCACHE set?", err)
	}
	// i USE SYMBOLS

	// create the interpreter
	i := interp.New(interp.Options{GoPath: os.Getenv("GOPATH"), GoCache: goCache})
	if _, err := i.Eval(source); err != nil {
		return nil, fmt.Errorf("An error occurred loading the template file: %v\n%v", absfilepath, err)
	}
	return i, nil
}
