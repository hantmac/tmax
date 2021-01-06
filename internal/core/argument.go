package core

import (
	"github.com/c-bata/go-prompt"
	"github.com/google/martian/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"tmax/setting"
)

type Argument map[string]string
type OutMap map[string][]map[string]string

var Args Argument

func TransYamlToOutMap(fileName string) OutMap {

	m := OutMap{}
	if !ExistFile(fileName) {
		log.Errorf("the .tmax.yaml not exist, please use `tmax generate` to get it ")
		return m
	}
	f, err := ioutil.ReadFile(fileName)
	err = yaml.Unmarshal(f, &m)
	if err != nil {
		log.Errorf("error: %v", err)
	}

	return m
}

func TransferYamlToMap(fileName string) Argument {
	args := Argument{}

	m := TransYamlToOutMap(fileName)

	for _, v := range m {
		for _, value := range v {
			for kk, vv := range value {
				args[kk] = vv
			}
		}
	}

	return args
}

func TransMapToYaml(om OutMap) ([]byte, error) {
	b, err := yaml.Marshal(om)
	if err != nil {
		log.Errorf("transfer map to yaml failed: %v", err)
		return b, err
	}
	return b, nil
}

func (o OutMap) InsertToMap(key string, m map[string]string) {

	if v, ok := o[key]; ok {
		v = append(v, m)
		o[key] = v
		return
	}

	value := make([]map[string]string, 0)
	value = append(value, m)
	o[key] = value
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

	args := TransferYamlToMap(setting.FileName)
	for k, v := range args {
		s = append(s, prompt.Suggest{Text: k, Description: v})
	}

	s = append(s, t...)

	return prompt.FilterFuzzy(s, d.GetWordBeforeCursor(), true)
}
