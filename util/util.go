package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
)

//出错拦截
func Catch() {
	if r := recover(); r != nil {
		for skip := 0; ; skip++ {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				break
			}
			go log.Printf("%v,%v\n", file, line)
		}
	}
}

//对象数组转成map数组
func ObjArrToMapArr(old []interface{}) []map[string]interface{} {
	if old != nil {
		new := make([]map[string]interface{}, len(old))
		for i, v := range old {
			new[i] = v.(map[string]interface{})
		}
		return new
	} else {
		return nil
	}
}

func IntAll(num interface{}) int {
	if i, ok := num.(int); ok {
		return int(i)
	} else if i0, ok0 := num.(int32); ok0 {
		return int(i0)
	} else if i1, ok1 := num.(float64); ok1 {
		return int(i1)
	} else if i2, ok2 := num.(int64); ok2 {
		return int(i2)
	} else if i3, ok3 := num.(float32); ok3 {
		return int(i3)
	} else if i4, ok4 := num.(string); ok4 {
		in, _ := strconv.Atoi(i4)
		return int(in)
	} else if i5, ok5 := num.(int16); ok5 {
		return int(i5)
	} else if i6, ok6 := num.(int8); ok6 {
		return int(i6)
	} else {
		return 0
	}
}

//读取配置文件
func ReadConfig(config ...interface{}) {
	var r *os.File
	filepath := "./config.json"
	if len(config) > 1 {
		filepath, _ = config[0].(string)
	}
	r, _ = os.Open(filepath)
	defer r.Close()
	bs, _ := ioutil.ReadAll(r)
	json.Unmarshal(bs, config[0])
}
