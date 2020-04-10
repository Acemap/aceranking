package cache

import (
	"bytes"
	"fmt"
	goCache "github.com/patrickmn/go-cache"
	"reflect"
	"runtime"
	"time"
)

var cacheVar *goCache.Cache

func init() {
	// Cache will expire at 5 days after created.
	cacheVar = goCache.New(5*24*time.Hour, 10*time.Minute)
}

// Wrapper for functions which need to be cached.
func Cache(in ...interface{}) interface{} {
	var keyBuf bytes.Buffer

	fun := reflect.ValueOf(in[0])
	// function name
	keyBuf.WriteString(runtime.FuncForPC(fun.Pointer()).Name())

	var args []reflect.Value
	for _, arg := range in[1:] {
		args = append(args, reflect.ValueOf(arg))
		keyBuf.WriteString(fmt.Sprintf("/%v", arg))
	}

	key := keyBuf.String()
	if value, ok := cacheVar.Get(key); ok {
		return value
	} else {
		ret := fun.Call(args)
		// For simplicity, only support function which only return one result.
		v := ret[0].Interface()
		cacheVar.SetDefault(key, v)
		return v
	}
}

//type OutputPair struct {
//	Key interface{}
//	Value interface{}
//}

// Wrapper for functions which need to be cached.
// List to list version
//func CacheList(in ...interface{}) []OutputPair {
//	var result []OutputPair
//	var keyBuf bytes.Buffer
//
//	fun := reflect.ValueOf(in[0])
//	// function name
//	keyBuf.WriteString(runtime.FuncForPC(fun.Pointer()).Name())
//
//	list := reflect.ValueOf(in[1])
//
//	var leftArgs []reflect.Value
//	for _, arg := range in[1:] {
//		leftArgs = append(leftArgs, reflect.ValueOf(arg))
//		keyBuf.WriteString(fmt.Sprintf("/%v", arg))
//	}
//
//	var remain []reflect.Value
//	l := list.Len()
//	for i := 0; i < l; i++ {
//		input := list.Index(i)
//		cacheKey := keyBuf.String() + fmt.Sprintf("/%v", value)
//		if output, ok := cacheVar.Get(cacheKey); ok {
//			result = append(result, OutputPair{input, output})
//		} else {
//			remain = append(remain, input)
//		}
//	}
//	if remain == nil {
//		return result
//	}
//
//	var args []reflect.Value
//	args = append(args, reflect.ValueOf(remain))
//	args = append(args, leftArgs...)
//	ret := fun.Call(args)[0]
//
//	rl := ret.Len()
//	for i := 0; i < rl; i++ {
//
//	}
//
//	key := keyBuf.String()
//	if value, ok := cacheVar.Get(key); ok {
//		return value
//	} else {
//		ret := fun.Call(leftArgs)
//		// For simplicity, only support function which only return one result.
//		v := ret[0].Interface()
//		cacheVar.SetDefault(key, v)
//		return v
//	}
//}

func Set(key string, value interface{}) {
	cacheVar.SetDefault(key, value)
}

func Get(key string) (interface{}, bool) {
	return cacheVar.Get(key)
}
