package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/stevemurr/cbot/asset"
)

type Config struct {
	AssetMap map[string]asset.Asset
	Assets   []string `json:"assets"`
}

func (c *Config) createAssetMap() {
	assetMap := map[string]asset.Asset{}
	for _, item := range c.Assets {
		assetMap[item] = asset.Asset{Name: item}
	}
	c.AssetMap = assetMap
}

func ReadConfig(filename string) (*Config, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var c Config
	if err := json.Unmarshal(f, &c); err != nil {
		return nil, err
	}

	c.createAssetMap()
	return &c, nil
}
