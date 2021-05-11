package config

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	viperIns *viper.Viper
)

func init() {
	viperIns = viper.New()
}

type Config interface {
	Init() error
}

func SetConfigFile(in string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	viperIns.SetConfigFile(in)
	err := viperIns.ReadInConfig()
	if err != nil {
		log.Println(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
}

func Unmarshal(cfg interface{}) error {
	err := viperIns.Unmarshal(cfg)
	if err != nil {
		return errors.Wrap(err, "Unmarshal error")
	}
	return nil
}

func UnmarshalKey(key string, cfg interface{}) error {
	err := viperIns.UnmarshalKey(key, cfg)
	if err != nil {
		return errors.Wrap(err, "UnmarshalKey error")
	}
	return nil
}

func LoadConfig(key string, cfg Config) error {
	err := viperIns.UnmarshalKey(key, cfg)
	if err != nil {
		return errors.Wrap(err, "LoadConfig error")
	}

	err = cfg.Init()
	if err != nil {
		return errors.Wrap(err, "Init error")
	}

	return nil
}
