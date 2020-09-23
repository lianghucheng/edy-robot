package util

import (
	"encoding/json"
	"fmt"
	"github.com/name5566/leaf/log"
)

func PrintObject(data interface{}) {
	m := map[string]interface{}{}
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err.Error())
		return
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Error(err.Error())
		return
	}
	for k, v := range m {
		fmt.Println(k, ": ", v)
	}
}
