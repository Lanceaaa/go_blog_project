package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// 初始化项目配置的基础属性 配置文件名、配置路径、配置类型
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	// 根据其默认值判断配置文件是否存在。若存在，则对读取配置的路径进行变更
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
