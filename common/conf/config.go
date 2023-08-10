package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config 函数入口
type Config struct {
}

// Init 初始化
func Init() {
	c := newConfig()
	if err := c.getConfig(); err != nil {
		log.Fatalf("load config err %s", err.Error())
	}
}

// newConfig new
func newConfig() *Config {
	return &Config{}
}

// getConfig 读取配置
func (c *Config) getConfig() error {
	conf := []eConfig{}
	conf = append(conf, eConfig{FileName: "app.toml", Value: &AppConf})
	return c.parse(conf)
}

type eConfig struct {
	FileName string
	Value    interface{}
}

// parse 解析配置
func (c *Config) parse(confs []eConfig) (err error) {
	for _, conf := range confs {
		_, err = toml.DecodeFile(checkFile("config/"+conf.FileName), conf.Value)
		if err != nil {
			return
		}
	}
	return
}

// checkFile 单元测试便利
func checkFile(file string) string {
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(file); err == nil {
			break
		} else {
			file = "../" + file
		}
	}
	return file
}
