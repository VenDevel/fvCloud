package utils

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var _dataCenter *dataCenter

var mutex sync.Mutex

func init() {
	newDataCenter()
	go task()
}

func task() {
	for {
		log.Println("start utils dataCenter task")

		time.Sleep(20 * time.Second) //10秒钟让放出去执行任务
	}
}

type dataCenter struct {
	data map[string]interface{}
	info map[string]*dataCenterKeyInfo
}

type dataCenterKeyInfo struct {
	Key     string
	gtime   int64
	addtime int64
}

func newDataCenter() *dataCenter {
	_dataCenter = &dataCenter{}
	_dataCenter.data = make(map[string]interface{})
	_dataCenter.info = make(map[string]*dataCenterKeyInfo)
	return _dataCenter
}

func GetInstance() *dataCenter {

	mutex.Lock()
	defer mutex.Unlock()
	if _dataCenter == nil {
		return newDataCenter()
	}
	return _dataCenter
}

func (this *dataCenter) Add(key string, value interface{}) error { //增加因此如果key存在会报错
	mutex.Lock()
	defer mutex.Unlock()
	log.Println("-------------------------------------")
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
	mutex.Lock()
	defer mutex.Unlock()
	value, ok := this.data[key]
	return value, ok
}

func (this *dataCenter) Set(key string, value interface{}) { //set 如果没有就会添加 如果有就会覆盖
	mutex.Lock()
	defer mutex.Unlock()
	this.data[key] = value
}

func (this *dataCenter) SetKeyAndInfo(key string, value interface{}, g int64) {
	mutex.Lock()
	defer mutex.Unlock()
	this.data[key] = value
	nowtime := time.Now().UnixNano()
	this.info[key] = &dataCenterKeyInfo{}
	this.info[key].Key = key
	this.info[key].addtime = nowtime
	this.info[key].gtime = nowtime + g
}

func (this *dataCenter) isExpired(key string) (bool, error) {
	info, ok := this.info[key]
	if !ok {
		return false, fmt.Errorf("not key")
	}
	return info.gtime <= time.Now().UnixNano(), nil
}
