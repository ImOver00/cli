package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	toml2 "github.com/BurntSushi/toml"
	"github.com/komkom/toml"
	"io/ioutil"
	"os"
)

// Config contains the basic TOML structure acceptable by sampgo.
type Config struct {
	Global struct {
		Sampctl bool `json:"sampctl"`
	} `json:"global"`

	Author struct {
		User string `json:"user"`
		Repo string `json:"repo"`
	} `json:"author"`

	Package struct {
		Name   string `json:"name"`
		Input  string `json:"input"`
		Output string `json:"output"`
	} `json:"package"`
}

// WriteToml allows you to write a toml file.
func WriteToml(fileName string, config Config) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	if err := toml2.NewEncoder(f).Encode(config); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

// ParseToml parses the toml file specified by the user.
func (p Parser) ParseToml() (Config, error) {
	if p.Dialect != Toml {
		return Config{}, fmt.Errorf("instance of parser is not using toml")
	}

	data, err := ioutil.ReadFile(p.FileName)
	if err != nil {
		return Config{}, fmt.Errorf("invalid file name")
	}

	// Our toml library leverages the default JSON decoder.
	decoder := json.NewDecoder(toml.New(bytes.NewBuffer(data)))

	var config Config

	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
