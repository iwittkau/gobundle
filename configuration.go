// Package gobundle implements data structures an install functions
package gobundle

import (
	"fmt"
	"io"
	"os/exec"
)

// Configuration is the project configuration struct
type Configuration []Package

// Package is a configuration element
type Package struct {
	Package string `json:"package"`
	Version string `json:"version"`
}

// DefaultConfiguration for gobundle initialization
var DefaultConfiguration = Configuration{
	{Package: "github.com/rakyll/govalidate", Version: "latest"},
	{Package: "github.com/google/gops", Version: "latest"},
	{Package: "honnef.co/go/tools/cmd/staticcheck", Version: "2020.1.3"},
}

// Install installs all packages of a configuration
func (c Configuration) Install(out io.Writer) error {
	for _, conf := range c {

		cmd := exec.Command("go", "get", fmt.Sprintf("%s@%s", conf.Package, conf.Version))
		cmd.Stdout = out
		cmd.Stderr = out
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
