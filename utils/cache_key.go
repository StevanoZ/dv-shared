package shrd_utils

import (
	"fmt"
	"reflect"
)

func BuildCacheKey(key string, identifier string, funcName string, args ...any) string {
	cacheKey := fmt.Sprintf("%s-%s-%s", key, identifier, funcName)
	cacheArgs := ""
	for index, arg := range args {

		if cacheArgs != "" {
			cacheArgs += "|"
		}

		v := reflect.ValueOf(arg)
		for i := 0; i < v.NumField(); i++ {
			if cacheArgs == "" {
				cacheArgs += fmt.Sprintf("%s:%v,", v.Type().Field(i).Name, v.Field(i).Interface())
			} else if index > 0 {
				cacheArgs += fmt.Sprintf("%s:%v,", v.Type().Field(i).Name, v.Field(i).Interface())
			} else {
				cacheArgs += fmt.Sprintf("%s:%v", v.Type().Field(i).Name, v.Field(i).Interface())
			}
		}
	}

	cacheKeyResult := fmt.Sprintf("%s|%s", cacheKey, cacheArgs)
	return string(cacheKeyResult[0 : len(cacheKeyResult)-1])
}

func BuildPrefixKey(keys ...string) string {
	prefixKey := ""

	for _, key := range keys {
		if prefixKey == "" {
			prefixKey += key
		} else {
			prefixKey += fmt.Sprintf("-%s", key)
		}
	}

	return prefixKey
}
