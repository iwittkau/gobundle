package gobundle_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/iwittkau/gobundle"
)

type noopWriter struct{}

func (w noopWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
func TestDump(t *testing.T) {
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	// defer w.Flush()
	conf, err := gobundle.Dump(noopWriter{})
	if err != nil {
		t.Error(err.Error())
	}
	data, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		t.Error(err.Error())
	}
	err = ioutil.WriteFile("examples/dumped.json", data, os.ModePerm)
	if err != nil {
		t.Error(err.Error())
	}
}
