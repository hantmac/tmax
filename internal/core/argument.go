package core

import (
	"github.com/c-bata/go-prompt"
	"github.com/google/martian/log"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type Argument map[string]string

var Args Argument

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
		for k, value := range v {
			args[k] = value
		}
	}

	return args
}

func ExistFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}

func Complete(d prompt.Document) []prompt.Suggest {

	var s []prompt.Suggest

	t := []prompt.Suggest{
		{Text: "exit", Description: "exit"},
	}

	args := TransferYamlToMap("~/.tmax.yaml")
	for k, v := range args {
		s = append(s, prompt.Suggest{Text: k, Description: v})
	}

	s = append(s, t...)

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
