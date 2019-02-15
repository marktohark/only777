package config

import "strings"

var cfgStore map[string]string

func Init() error {
	cfgMap, err := loadCfg()
	if err != nil {
		return err
	}
	cfgStore = *cfgMap
	return nil
}

func Get(key string) (string, bool) {
	str, ok := cfgStore[strings.ToUpper(key)]
	return str, ok
}

func Set(key string, val string) error {
	cfgStore[strings.ToUpper(key)] = val
	err := saveCfg(&cfgStore)
	if err != nil {
		return err
	}
	return nil
}