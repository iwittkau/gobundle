package gobundle

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Dump tries to dump currently installed tools and version from $GOPATH/bin
func Dump(out io.Writer) (Configuration, error) {
	conf := Configuration{}
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return conf, errors.New("GOPATH not set, no binaries to analyze")
	}
	binDir := strings.Join([]string{goPath, "bin"}, string(os.PathSeparator))
	files, err := ioutil.ReadDir(binDir)
	if err != nil {
		return conf, err
	}
	skipped := 0
	fmt.Fprintln(out, "Binary\tPackage\tVersion")
	fmt.Fprintln(out, "------\t----\t-------")
	for i := range files {
		if files[i].IsDir() {
			fmt.Fprintln(out, files[i], "is a directory")
			continue
		}
		absFilePath := strings.Join([]string{binDir, files[i].Name()}, string(os.PathSeparator))

		// a better regex would be helpful here ...
		cmd := exec.Command("go", "tool", "objdump", "-s", "main.main$", absFilePath)
		output, _ := cmd.CombinedOutput()

		// extract first line of output
		path := strings.Split(string(output), "\n")[0]

		// replace irrelevant strings
		path = strings.ReplaceAll(path, "TEXT main.main(SB) ", "")

		// check if installed via `go get`
		if !strings.Contains(path, "@") {
			// probably built locally
			fmt.Fprintf(out, "* %s\t%s\t[skipped]\n", files[i].Name(), path)
			skipped++
			continue
		}

		// extract version
		version := strings.Split(path, "@")[1]
		version = strings.Split(version, "/")[0]

		// clean-up path
		path = strings.ReplaceAll(path, version, "")
		path = strings.ReplaceAll(path, "@", "")
		path = strings.Split(path, "pkg/mod/")[1]

		// strip any .go files
		path = filepath.Dir(path)

		fmt.Fprintf(out, "%s\t%s\t%s\n", files[i].Name(), path, version)
		conf = append(conf, Package{Package: path, Version: version})

	}
	if skipped > 0 {
		bin := "binary"
		pron := "it was"
		if skipped > 1 {
			bin = "binaries"
			pron = "they were"
		}
		fmt.Fprintf(out, "\nINFO: skipped %d %s because %s built locally\n", skipped, bin, pron)
	}
	return conf, nil
}
