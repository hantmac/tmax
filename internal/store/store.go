package store

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cast"
	"gopkg.in/yaml.v2"

	"github.com/hantmac/tmax/setting"
)

const defaultGroupName = "default"

type Shortcuts map[string]string

type store struct {
	groups []string // TODO: unused
	shortcuts Shortcuts
	groupedShortcuts map[string]interface{}
}

func Store() *store {
	if !Exists(setting.ConfigPath) {
		fmt.Println("tmax.yaml not exists, use `tmax generate` to generate it first")
		os.Exit(1)
	}

	shortcuts, err := ReadConfig(setting.ConfigPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	groupedShortcuts, err := getConfig(setting.ConfigPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &store{
		shortcuts: shortcuts,
		groupedShortcuts: groupedShortcuts,
	}
}

func (s *store) GetFullCommand(short string) (string, bool) {
	long, ok  := s.shortcuts[short]
	return long, ok
}

func (s *store) Shortcuts() Shortcuts {
	return s.shortcuts
}

func (s *store) AddRecord(key, value string) error {
	ks := strings.Split(key, ".")
	group := ks[0]
	if len(ks) == 1 {
		group = defaultGroupName
	}
	s.addRecordToGroup(ks[len(ks)-1], value, group)

	data, err := yaml.Marshal(s.groupedShortcuts)
	if err != nil {
		fmt.Printf("can not save new command: %s\n", err)
	}

	return ioutil.WriteFile(setting.ConfigPath, data, os.ModePerm)
}

func (s *store) addRecordToGroup(key, value, group string) {
	g, ok := s.groupedShortcuts[group]
	if !ok {
		g = make(map[string]interface{})
	}

	switch val := g.(type) {
	case map[interface{}]interface{}:
		sm := cast.ToStringMap(val)
		sm[key] = value
		s.groupedShortcuts[group] = sm
	case map[string]interface{}:
		val[key] = value
		s.groupedShortcuts[group] = val
	}
}

func ReadConfig(fileName string) (Shortcuts, error) {
	m, err := getConfig(fileName)
	if err != nil {
		return nil, err
	}

	flattened := Flatten(m)
	shortenKeys(flattened)

	return flattened, nil
}

func getConfig(fileName string) (map[string]interface{}, error) {
	if !Exists(fileName) {
		return nil, fmt.Errorf("tmax.yaml not exists, use `tmax generate` to generate it first")
	}

	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal(f, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func flatten(key string, value interface{}, store map[string]string) {
	if key == "" {
		m, ok := value.(map[string]interface{})
		if !ok {
			return
		}
		for k, v := range m {
			flatten(k, v, store)
		}

		return
	}

	switch val := value.(type) {
	case string, int, bool, float32, float64, int32, int64:
		store[key] = cast.ToString(val)
	case map[interface{}]interface{}:
		sm := cast.ToStringMap(val)
		for k, v := range sm {
			flatten(fmt.Sprintf("%s.%s", key, k), v, store)
		}
	case map[string]interface{}:
		for k, v := range val {
			flatten(fmt.Sprintf("%s.%s", key, k), v, store)
		}
	default:
		fmt.Printf("can not parse value %v (type %T)", val, val)
	}
}

// Flatten converts a nested dict like {a: {b: c}, d: e} to a flatten one {a.b: c, d: e}
func Flatten(data map[string]interface{}) map[string]string {
	seed := make(map[string]string)
	flatten("", data, seed)

	return seed
}

func shortenKeys(data map[string]string) {
	for k, v := range data {
		ks := strings.Split(k, ".")
		if len(ks) > 1 {
			data[ks[len(ks)-1]] = v
		}
	}
}


func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}
