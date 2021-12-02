package sf

import (
	"encoding/base64"
	"fmt"
	"sync"

	sfg "golang.org/x/sync/singleflight"
)

var sfForAll = &sfg.Group{}
var sfGroupKeyMap = make(map[string]*sfg.Group)
var sfKeyMapMutex = sync.RWMutex{}

func ByKey(key string, fn func() (interface{}, error)) (res interface{}, shared bool, err error) {
	var sfgroup *sfg.Group
	hasGroup := false

	sfKeyMapMutex.RLock()
	sfgroup, hasGroup = sfGroupKeyMap[key]
	sfKeyMapMutex.RUnlock()

	if !hasGroup || sfgroup == nil {
		sfgroup = getGroup(key)
	}

	res, err, shared = sfgroup.Do(key, fn)
	return
}

func ByParams(params []interface{}, fn func() (interface{}, error)) (res interface{}, shared bool, err error) {
	return ByParamsWithPrefix(params, "", fn)
}

func ByParamsWithPrefix(params []interface{}, keyPrefix string, fn func() (interface{}, error)) (res interface{}, shared bool, err error) {
	rawKey := ""
	for i := 0; i < len(params); i++ {
		rawKey = rawKey + "|" + fmt.Sprintf("%+v", params[i])
	}
	key := keyPrefix + ":" + base64.StdEncoding.EncodeToString([]byte(rawKey))
	return ByKey(key, fn)
}

func getGroup(key string) *sfg.Group {
	gRaw, _, _ := sfForAll.Do(key, func() (interface{}, error) {
		sfgroup := &sfg.Group{}
		sfKeyMapMutex.Lock()
		sfGroupKeyMap[key] = sfgroup
		sfKeyMapMutex.Unlock()
		return sfgroup, nil
	})
	return gRaw.(*sfg.Group)
}
