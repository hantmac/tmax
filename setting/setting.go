package setting

import (
	"github.com/mitchellh/go-homedir"
	"path"
)

var (
	homedirStr, _ = homedir.Dir()
 	ConfigPath    = path.Join(homedirStr, ".tmax.yaml")
)