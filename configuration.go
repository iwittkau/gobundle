package gobundle

import (
	"fmt"
	"os"
	"os/exec"
)

// Configuration is the project configuration struct
type Configuration []Package

// Package is a configuration element
type Package struct {
	Package string `json:"package"`
	Version string `json:"version"`
}

// Install installs all packages of a configuration
func (c Configuration) Install() error {
	for _, conf := range c {

		cmd := exec.Command("go", "get", fmt.Sprintf("%s@%s", conf.Package, conf.Version))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
