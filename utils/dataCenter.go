package utils

import (
	"fmt"
	"log"
	"sync"
)

var _dataCenter *dataCenter

var mutex sync.Mutex

func init() {
	newDataCenter()
}

type dataCenter struct {
	data *map[string]interface{}
	info *map[string]keyinfo
}

type keyinfo struct {
	Key     string
	gtime   int64
	addtime int64
}

func newDataCenter() *dataCenter {
	_dataCenter = make(dataCenter)
	_dataCenter.data = make(map[string]interface{})
	_dataCenter.info = make(keyinfo)
	return _dataCenter
}

func GetInstance() {
	mutex.Lock()
	defer mutex.Unlock()
	if _dataCenter == nil {
		newDataCenter()
	}
}

func (this *dataCenter) Add(key string, value interface{}) error { //增加因此如果key存在会报错
	mutex.Lock()
	defer mutex.Unlock()
	_, ok := this.data[key]
	if ok {
		return fmt.Errorf("key repeat")
	}
	this.data[key] = value
	return nil
}

func (this *dataCenter) Delete(key string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(this.data, key)
}

func (this *dataCenter) Get(key string) (interface{}, bool) {
	return this.data[key]
}

func (this *dataCenter) Set(key string, value interface{}) { //set 如果没有就会添加 如果有就会覆盖
	this.data[key] = value
}
