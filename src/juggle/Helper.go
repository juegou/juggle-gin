// +----------------------------------------------------------------------
// | Juggle [ 让我们能更好的杂耍 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2020 http://www.XXXXXX.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: dingo <djhui1987@gmail.com>
// +----------------------------------------------------------------------

// 辅助函数库
package juggle

import (
	"io/ioutil"
	"os"
)

// 加载根目录下的配置文件application.yaml
func LoadConfigFile() []byte {
	dir,err := os.Getwd() // 获取当前目录
	if err != nil {
		Error(err)
	}
	//log.Println(dir + "/app/application.yaml")
	b,err := ioutil.ReadFile(dir + "/app/application.yaml")
	if err != nil {
		return nil
	}
	return b
}
