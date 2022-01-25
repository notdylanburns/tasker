package config

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Config struct {
	port uint16
}

const ConfigurationDir = "tasker"
const ConfigurationFile = "taskerd.conf"

const DefaultPort = uint16(43976)

func New() *Config {
	cfg := new(Config)
	cfg.port = DefaultPort

	return cfg
}

func Exists(dir string) bool {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func Load(file string) (*Config, error) {
	return nil, nil
}

func Create(dir string) (*Config, error) {
	configDir := path.Join(dir, ConfigurationDir)
	configFile := path.Join(configDir, ConfigurationFile)

	var err error

	if !Exists(configDir) {
		err = os.Mkdir(configDir, 0777)
		if err != nil {
			return nil, err
		}
	}

	if !Exists(configFile) {
		f, err := os.Create(configFile)
		if err != nil {
			return nil, err
		}

		defer func(f *os.File) {
			err = f.Close()
		}(f)

		cfg := New()
		err = cfg.Write(f)
		if err != nil {
			return nil, err
		}

		return cfg, err

	} else {
		return nil, os.ErrExist
	}
}

func (cfg *Config) Write(f *os.File) error {
	s := fmt.Sprintf("%+v", cfg)
	s = strings.Replace(s, "{", "", -1)
	s = strings.Replace(s, "}", "", -1)
	s = strings.Replace(s, ":", "=", -1)
	_, e := fmt.Fprintf(f, "[taskerd.conf]\n%s", strings.Join(strings.Split(s, ", "), "\n"))

	return e
}

func (cfg *Config) WriteFile(f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err = file.Close()
	}(file)

	err = cfg.Write(file)
	if err != nil {
		return err
	}

	return err
}
