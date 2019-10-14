package gobundle_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/iwittkau/gobundle"
)

func TestConfiguration_MarshalJSON(t *testing.T) {
	c := gobundle.Configuration{
		{"github.com/iwittkau/gobundle/cmd/gobundle", "latest"},
	}

	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		t.Error(err.Error())
	}
	err = ioutil.WriteFile("examples/gobundle.json", data, os.ModePerm)

	d, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}
	d = path.Join(d, "_test_gopath")
	os.Setenv("GOPATH", d)
	err = c.Install()
	if err != nil {
		t.Error(err.Error())
	}

}
