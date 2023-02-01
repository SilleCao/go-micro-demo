package common

import (
	"encoding/json"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/schema"
)

func GetReqIdKey() string {
	return string(REQ_ID)
}

func GetReqIdValue(ctx *gin.Context) string {
	reqId := ctx.Value(REQ_ID)
	if reqId == nil {
		reqId = ""
	}
	return reqId.(string)
}

func GetReqBodyAsMap(ctx *gin.Context) (map[string]interface{}, error) {
	bodyData, _ := io.ReadAll(ctx.Request.Body)
	data := make(map[string]interface{})
	err := json.Unmarshal(bodyData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ConvertToGormKey(m map[string]interface{}, obj any) {
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		jsonKey := sf.Tag.Get("json")
		tagSetting := schema.ParseTagSetting(sf.Tag.Get("gorm"), ";")
		val, ok := m[jsonKey]
		if ok {
			delete(m, jsonKey)
			m[tagSetting["COLUMN"]] = reflect.ValueOf(val).Convert(sf.Type).Interface()
		}
	}
}
