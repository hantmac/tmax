package setting

import (
	"path"

	"github.com/mitchellh/go-homedir"
)

var (
	homedirStr, _ = homedir.Dir()
 	ConfigPath    = path.Join(homedirStr, ".tmax.yaml")
)
