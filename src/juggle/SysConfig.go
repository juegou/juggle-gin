// +----------------------------------------------------------------------
// | Juggle [ 让我们能更好的杂耍 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2020 http://www.XXXXXX.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: dingo <djhui1987@gmail.com>
// +----------------------------------------------------------------------

// 使用yaml解析配置文件
package juggle

import (
	"gopkg.in/yaml.v2"
	"log"
)

// 服务配置
type ServerConfig struct {
	Port int // 端口
}

// 用户自定义配置
type UserConfig map[interface{}]interface{}


// 系统配置
type SystemConfig struct {
	Server *ServerConfig
	Config UserConfig
}

// 初始化系统配置
func NewSystemConfig() *SystemConfig {
	return &SystemConfig{Server: &ServerConfig{Port:8080}} // 默认端口为8080
}

// 加载yaml配置文件并解析
func InitConfig() *SystemConfig {
	sysConfig := NewSystemConfig()
	if b := LoadConfigFile(); b != nil {
		err := yaml.Unmarshal(b,sysConfig)
		if err != nil {
			log.Fatal(err)
		}
	}
	return sysConfig
}

