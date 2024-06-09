package zerome

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type QuerierConfig struct {
	Address string            `yaml:"address"`
	Headers map[string]string `yaml:"headers"`
	Type    string            `yaml:"type"`
}

type WriterConfig struct {
	Address string            `yaml:"address"`
	Headers map[string]string `yaml:"headers"`
	Timeout time.Duration     `yaml:"timeout"`
	Type    string            `yaml:"type"`
}

type Config struct {
	Queriers map[string]QuerierConfig `yaml:"queriers"`
	Writers  map[string]WriterConfig  `yaml:"writers"`
	Metrics  []Metric                 `yaml:"metrics"`
}

func LoadConfig(configPath string, cfg *Config) error {
	configFile, err := os.Open(configPath)
	if err != nil {
		return err
	}

	return yaml.NewDecoder(configFile).Decode(cfg)
}
