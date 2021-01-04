package core

import (
	"github.com/google/martian/log"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type Argument map[string]string

func TransferYamlToMap(fileName string) Argument {
	args := Argument{}
	m := make(map[string]map[string]string)
	homedirStr, err := homedir.Dir()
	if err != nil {
		log.Errorf("get home dir failed: %v", err)
	}
	fileName = path.Join(homedirStr, ".tmax.yaml")
	if !ExistFile(fileName) {
		log.Errorf("the .tmax.yaml not exist, please use `tmax generate` to get it ")
		return args
	}
	f, err := ioutil.ReadFile(fileName)
	err = yaml.Unmarshal(f, &m)
	if err != nil {
		log.Errorf("error: %v", err)
	}

	for _, v := range m {
		for k, value := range  v {
			args[k] = value
		}
	}

	return args
}

func ExistFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}
