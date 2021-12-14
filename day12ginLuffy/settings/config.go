package settings

import (
	"encoding/json"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/confmap"
	"log"
)

type Config struct {
	ServerConfig ServerConfig `json:"server"`
	DBConfig DBConfig `json:"db"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type DBConfig struct {
	DataSourceName string `json:"dataSourceName"`
}

func (dbConfig *DBConfig) MarshalJSON() ([]byte, error)  {
	m := map[string]interface{}{
		"dataSourceName": "[PROTECTED]",
	}
	return json.Marshal(m)
}

// load from default config
func Load()(*Config, error)  {
	k := koanf.New(".")
	err := k.Load(confmap.Provider(defaultConfig,"."), nil)
	if err != nil{
		log.Printf("fail to load default config, err：%v",err)
	}

	var cfg Config
	if err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{
		Tag: "json",
		FlatPaths: false,
	}); err != nil{
		log.Printf("failed to unmarshal with conf. err: %v", err)
		return nil, err
	}
	return &cfg, err
}

