package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var configLog = log.WithField("from", "main")

type CliConfig struct {
	Endpoint     string `yaml:"endpoint"`
	AuthEndpoint string `yaml:"auth-endpoint"`
	Token        string `yaml:"token"`
	UserId       string `yaml:"user-id"`
}

type ConfigurationManager struct {
	FilePath string
}

func NewConfigurationManager(fileLocation string) (*ConfigurationManager, error) {
	fileLocation = os.ExpandEnv(fileLocation)

	configFile, err := filepath.Abs(fileLocation)
	if err != nil {
		return nil, errors.Wrapf(err,
			"Couldn't get absolute path of config file %s",
			fileLocation)
	}

	if _, err := os.Stat(configFile); err != nil {
		if !os.IsNotExist(err) {
			return nil, errors.Wrapf(err,
				"Couldn't get file information for file %s",
				configFile)
		}
	}

	return &ConfigurationManager{
		FilePath: configFile,
	}, nil
}

// Write writes the configuration passed as argument to the
// file (ConfigurationManager.FilePath).
func (cfg *ConfigurationManager) Write(config *CliConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrapf(err,
			"Couldn't marshal config %s",
			spew.Sdump(config))
	}

	return ioutil.WriteFile(cfg.FilePath, data, 0644)
}

// Read reads the configuration indicated by FilePath and
// returns it in the form of `CliConfig` object.
func (cfg *ConfigurationManager) Read() (*CliConfig, error) {
	var t = &CliConfig{}

	_, err := os.Stat(cfg.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return t, nil
		}
		return nil, errors.Wrapf(err,
			"Couldn't get file info for file %s",
			cfg.FilePath)
	}

	data, err := ioutil.ReadFile(cfg.FilePath)
	if err != nil {
		return nil, errors.Wrapf(err,
			"Couldn't read file %s", cfg.FilePath)
	}

	if err = yaml.Unmarshal(data, t); err != nil {
		return nil, errors.Wrapf(err,
			"Couldn't unmarshal data from file %s: %v",
			cfg.FilePath, string(data))
	}

	return t, nil
}
