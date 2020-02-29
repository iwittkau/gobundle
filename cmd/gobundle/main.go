package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime/debug"
	"text/tabwriter"

	"github.com/iwittkau/gobundle"
)

var version = "development"

const bundleFile = "gobundle.json"

func main() {
	log.SetFlags(0)
	var (
		file    string
		init    bool
		version bool
		dump    bool
	)

	flag.StringVar(&file, "f", "", "set configuation file")
	flag.BoolVar(&init, "i", false, "init configuration in home directory")
	flag.BoolVar(&version, "v", false, "print version")
	flag.BoolVar(&dump, "d", false, "dumps installed tools and versions to gobundle.json")
	flag.Parse()
	if version {
		println(buildInfoVersion())
		return
	}
	if dump {
		err := dumps()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if init {
		p := getDefaultConfigPath(home)
		_, err := os.Stat(p)
		if err == nil {
			log.Fatalf("%s already exists\n", p)
		}
		if err := initConfiguration(p); err != nil {
			log.Fatal(err)
		}
		log.Println("initialized configuration", p)
		os.Exit(0)
	}

	if file == "" {
		file = getDefaultConfigPath(home)
	}

	var (
		data []byte
		conf gobundle.Configuration
	)

	data, err = ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chdir(home)
	if err != nil {
		log.Fatal(err)
	}
	err = conf.Install(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("all go binaries successfully installed.")

}

func getDefaultConfigPath(homeDir string) string {

	return path.Join(homeDir, ".gobundle")
}

func initConfiguration(path string) error {

	data, err := json.MarshalIndent(gobundle.DefaultConfiguration, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, os.ModePerm)

}

func buildInfoVersion() string {
	info, ok := debug.ReadBuildInfo()
	switch {
	case !ok:
		return version
	case info.Main.Version == "(devel)":
		return version
	default:
		return info.Main.Version
	}
}

func dumps() error {
	_, err := os.Stat(bundleFile)
	if err == nil {
		return fmt.Errorf("%s already exists", bundleFile)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()
	conf, err := gobundle.Dump(w)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(bundleFile, data, os.ModePerm)
}
