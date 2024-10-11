package tools

import (
	"encoding/json"
	"go-gin-template/config/log"
)

// Json 根据一个context, 获取一个json字符串
//
// @Author: 罗德
// @Date: 2023/9/13
func Json(context interface{}) string {
	jsonByte, err := json.Marshal(context)
	if err != nil {
		log.Errorf("转换为JSON字符串失败: %s", err.Error())
		return ""
	}

	jsonString := string(jsonByte)
	if jsonString == "null" {
		return ""
	}
	return jsonString
}

// JsonFmt 根据一个context字符串, 获取一个json字符串带缩进
//
// @Author: 罗德
// @Date: 2023/9/13
func JsonFmt(context string) string {
	// 解析JSON字符串为一个空接口
	var jsonData interface{}
	err := json.Unmarshal([]byte(context), &jsonData)
	if err != nil {
		log.Errorf("解析JSON字符串失败: %s", err.Error())
		return ""
	}

	// 格式化JSON字符串为带缩进的JSON字符串
	indentedJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Errorf("格式化JSON字符串失败: %s", err.Error())
		return ""
	}

	return string(indentedJSON)
}

// JsonMap 根据一个json字符串, 获取一个map
func JsonMap(context string) map[string]interface{} {
	var data map[string]interface{}
	if len(context) <= 0 {
		return data
	}
	if err := json.Unmarshal([]byte(context), &data); err != nil {
		log.Errorf("解析JSON失败: %s", err.Error())
		return nil
	}
	return data
}

// JsonMapString 根据一个json字符串, 获取一个map[string]string
func JsonMapString(context string) map[string]string {
	var data map[string]string
	if len(context) <= 0 {
		return data
	}
	if err := json.Unmarshal([]byte(context), &data); err != nil {
		log.Errorf("解析JSON失败: %s", err.Error())
		return nil
	}
	return data
}

// JsonMapArray 根据一个json字符串, 获取一个 map[string]interface 数组
func JsonMapArray(context string) []map[string]interface{} {
	var data []map[string]interface{}
	if len(context) <= 0 {
		return data
	}
	if err := json.Unmarshal([]byte(context), &data); err != nil {
		log.Errorf("解析JSON失败: %s", err.Error())
		return nil
	}
	return data
}
