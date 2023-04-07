package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type RealTimeCfg struct {
	cfg *viper.Viper
}

var (
	rtc   = RealTimeCfg{}
	lock  = sync.RWMutex{}
	cache = make(map[string]interface{})

	svrCfg config
)

type config struct {
	Server map[string]map[string]interface{} `toml:"server"`
}

func Init(path string, cfg ...string) error {
	lock.Lock()
	defer lock.Unlock()
	v, err := loadConfig(path, cfg...)
	if err != nil {
		return err
	}
	rtc.cfg = v
	loadServerCfg()
	return nil
}

func loadServerCfg() {
	err := rtc.cfg.Unmarshal(&svrCfg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetSvrCfgMap() map[string]map[string]interface{} {
	return svrCfg.Server
}

func GetSvrCfg(id string) map[string]interface{} {
	cfg, ok := svrCfg.Server[id]
	if !ok {
		cfg = make(map[string]interface{})
	}
	return cfg
}
func loadConfig(path string, cfg ...string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(path) // 如果没有指定配置文件，则解析默认的配置文件
	for _, c := range cfg {
		v.SetConfigName(c)
		v.SetConfigType("toml") // 设置配置文件格式为toml
		if err := v.MergeInConfig(); err != nil {
			return nil, err
		}
	}
	return v, nil
}

func GetConfig() *viper.Viper {
	return rtc.cfg
}

func getFromCache(k string) (interface{}, bool) {
	lock.RLock()
	defer lock.RUnlock()
	ret, ok := cache[k]
	return ret, ok
}

func clearCache() {
	lock.Lock()
	defer lock.Unlock()
	cache = make(map[string]interface{})
}

func insertToCache(k string, v interface{}) {
	lock.Lock()
	defer lock.Unlock()
	cache[k] = v
}
func GetStringDef(k string, def string) string {
	if v, o := getFromCache(k); o {
		if rv, ok := v.(string); ok {
			return rv
		}
	}
	c := GetConfig()
	if c.IsSet(k) {
		v := c.GetString(k)
		insertToCache(k, v)
		return v
	} else {
		return def
	}
}
func GetString(k string) string {
	if v, o := getFromCache(k); o {
		if rv, ok := v.(string); ok {
			return rv
		}
	}
	c := GetConfig()
	v := c.GetString(k)
	insertToCache(k, v)
	return v
}
func GetIntDef(k string, def int) int {
	if v, o := getFromCache(k); o {
		if rv, ok := v.(int); ok {
			return rv
		}
	}
	c := GetConfig()
	if c.IsSet(k) {
		v := c.GetInt(k)
		insertToCache(k, v)
		return v
	} else {
		return def
	}
}
func GetInt(k string) int {
	if v, o := getFromCache(k); o {
		if rv, ok := v.(int); ok {
			return rv
		}
	}
	c := GetConfig()
	v := c.GetInt(k)
	insertToCache(k, v)
	return v
}
func GetBoolDef(k string, d bool) bool {
	if v, o := getFromCache(k); o {
		if rv, ok := v.(bool); ok {
			return rv
		}
	}
	c := GetConfig()
	if c.IsSet(k) {
		v := c.GetBool(k)
		insertToCache(k, v)
		return v
	} else {
		return d
	}
}
