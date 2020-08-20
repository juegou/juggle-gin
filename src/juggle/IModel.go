// +----------------------------------------------------------------------
// | Juggle [ 让我们能更好的杂耍 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2020 http://www.XXXXXX.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: dingo <djhui1987@gmail.com>
// +----------------------------------------------------------------------

// 用于规范Model代码
package juggle

import (
	"encoding/json"
	"log"
)

type IModel interface {
	String() string
}

type Models string

func MakeModels(v interface{}) Models {
	b,err := json.Marshal(v)
	if err != nil {
		log.Println(err.Error())
	}
	return Models(b)
}
