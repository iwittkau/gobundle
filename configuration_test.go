package gobundle_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/iwittkau/gobundle"
)

var c = gobundle.Configuration{
	{Package: "github.com/rakyll/govalidate", Version: "latest"},
	{Package: "github.com/google/gops", Version: "latest"},
	{Package: "honnef.co/go/tools/cmd/staticcheck", Version: "2020.1.3"},
}

func TestConfiguration_MarshalJSON(t *testing.T) {

	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		t.Error(err.Error())
	}
	err = ioutil.WriteFile("examples/gobundle.json", data, os.ModePerm)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestConfiguration_Install(t *testing.T) {
	tmp, err := ioutil.TempDir("", t.Name())
	if err != nil {
		t.Error(err.Error())
		return
	}
	conf := make(gobundle.Configuration, 1)
	conf[0] = c[0]

	oldenv := os.Getenv("GOPATH")
	defer func() {
		os.Setenv("GOPATH", oldenv)
	}()
	os.Setenv("GOPATH", tmp)
	err = conf.Install(noopWriter{})
	if err != nil {
		t.Error(err.Error())
	}
}
