package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/iwittkau/gobundle"
)

func main() {
	log.SetFlags(0)
	var (
		file string
		init bool
	)

	flag.StringVar(&file, "f", "", "set configuation file")
	flag.BoolVar(&init, "i", false, "init configuration in home directory")
	flag.Parse()

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if init {
		p := getDefaultConfigPath(home)
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
	json.Unmarshal(data, &conf)

	err = os.Chdir(home)
	if err != nil {
		log.Fatal(err)
	}
	err = conf.Install()
	if err != nil {
		log.Fatal(err)
	}
}

func getDefaultConfigPath(homeDir string) string {

	return path.Join(homeDir, ".gobundle")
}

func initConfiguration(path string) error {

	c := gobundle.Configuration{
		{Package: "github.com/iwittkau/gobundle/cmd/gobundle", Version: "latest"},
	}

	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, os.ModePerm)

}
