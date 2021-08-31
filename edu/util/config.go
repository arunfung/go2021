package util

import (
	"flag"
)

var (
	instance *uconfig
	conf     = flag.String("conf", "../etc/config.json", "cmdedu . config")
)
type uconfig struct {
	BasePath string `json:"base_path"`
	DataPath string `json:"data_path"`
}

func init() {
	flag.Parse()
	NewUConfigWithFile(*conf)
}

func NewUConfigWithFile(name string) *uconfig {
	if instance == nil {
		c := &uconfig{}
		ReadJosn(name, c)
		instance = c // <--- NOT THREAD SAFE
	}
	return instance
}

func GetConfig() *uconfig {
	return instance
}

func (u *uconfig) GetDataPath() string {
	return u.DataPath
}