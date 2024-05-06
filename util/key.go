/*
 * @Author: nijineko
 * @Date: 2024-04-22 10:46:20
 * @LastEditTime: 2024-04-22 10:54:58
 * @LastEditors: nijineko
 * @Description:
 * @FilePath: \gorm-cache\util\key.go
 */
package util

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

func GenInstanceId() string {
	charList := []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().Unix())
	length := 5
	str := make([]byte, 0)
	for i := 0; i < length; i++ {
		str = append(str, charList[rand.Intn(len(charList))])
	}
	return string(str)
}

func GenPrimaryCacheKey(KeyPrefix string, instanceId string, tableName string, primaryKey string) string {
	return fmt.Sprintf("%s%s:p:%s:%s", KeyPrefix, instanceId, tableName, primaryKey)
}

func GenPrimaryCachePrefix(KeyPrefix string, instanceId string, tableName string) string {
	return KeyPrefix + instanceId + ":p:" + tableName
}

func GenSearchCacheKey(KeyPrefix string, instanceId string, tableName string, sql string, vars ...interface{}) string {
	buf := strings.Builder{}
	buf.WriteString(sql)
	for _, v := range vars {
		pv := reflect.ValueOf(v)
		if pv.Kind() == reflect.Ptr {
			buf.WriteString(fmt.Sprintf(":%v", pv.Elem()))
		} else {
			buf.WriteString(fmt.Sprintf(":%v", v))
		}
	}
	return fmt.Sprintf("%s%s:s:%s:%s", KeyPrefix, instanceId, tableName, buf.String())
}

func GenSearchCachePrefix(KeyPrefix string, instanceId string, tableName string) string {
	return KeyPrefix + instanceId + ":s:" + tableName
}
