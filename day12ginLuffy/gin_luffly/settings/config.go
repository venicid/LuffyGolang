package settings

import (
	"encoding/json"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/confmap"
	"log"
)

type Config struct {
	ServerConfig ServerConfig `json:"server"`
	DBConfig     DBConfig     `json:"db"`
}

type ServerConfig struct {
	Port int `json: "port"`
}

type DBConfig struct {
	DataSourceName string `json: "dataSourceName"`
}

func (c *DBConfig) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"dataSourceName": "[PROTECTED]",
	}
	return json.Marshal(m)
}

func Load()(*Config, error){
	k := koanf.New(".")
	// load from default config
	err := k.Load(confmap.Provider(defalutConfig, "."), nil)
	if err != nil {
		log.Printf("faild to load default config, err: %v", err)
	}
	var cfg Config
	if err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err!=nil{
		log.Printf("failed to unmarshal with conf. err: %v", err)
		return nil, err
	}
	return &cfg, err
}
