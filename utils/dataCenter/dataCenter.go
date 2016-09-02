// @APIVersion 1.0.0
// @Title dataCenter pkg
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact Ven
package dataCenter

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var _dataCenter *dataPoll

var mutex sync.Mutex

var dataCenterClearCycle time.Duration = 30 * time.Second

func init() {
	GetInstance()
}

const (
	Nanosecond  int64 = 1                  //纳秒
	Microsecond       = 1000 * Nanosecond  //微妙
	Millisecond       = 1000 * Microsecond //毫秒
	Second            = 1000 * Millisecond // 秒
	Minute            = 60 * Second        //分
	Hour              = 60 * Minute        //时
	Day               = 24 * Hour          //天
)

// @Title 数据存放处
type dataPoll struct {
	data map[string]*dataCenterKeyInfo
}

// @Title dataCenter 定时任务
func (this *dataPoll) task() {
	for {
		log.Println("start utils dataCenter task")

		time.Sleep(dataCenterClearCycle) //10秒钟让放出去执行任务
		_dataCenter.clearDataCenter()
	}
}

type dataCenterKeyInfo struct {
	key     string
	gtime   int64
	addtime int64
	body    interface{}
}

func newDataCenter() *dataPoll {
	_dataCenter = &dataPoll{}
	_dataCenter.data = make(map[string]*dataCenterKeyInfo)
	return _dataCenter
}

func GetInstance() *dataPoll {
	mutex.Lock()
	defer mutex.Unlock()
	if _dataCenter == nil {
		go _dataCenter.task()
		return newDataCenter()
	}
	go _dataCenter.task()
	return _dataCenter
}

func (this *dataPoll) SetDataCenterClearCycle(dt time.Duration) *dataPoll {
	mutex.Lock()
	defer mutex.Unlock()
	dataCenterClearCycle = dt
	return this
}

func (this *dataPoll) Add(key string, value interface{}) error { //增加因此如果key存在会报错
	mutex.Lock()
	defer mutex.Unlock()
	_, ok := this.data[key]
	if ok {
		return fmt.Errorf("key repeat")
	}

	this.data[key] = &dataCenterKeyInfo{body: value, key: key, addtime: time.Now().UnixNano(), gtime: -1} //没有过期时间的默认-1标识
	return nil
}

func (this *dataPoll) Delete(key string) {
	mutex.Lock()
	defer mutex.Unlock()
	if !_dataCenter.isExpired(key) {
		delete(this.data, key)
	}
}

func (this *dataPoll) Get(key string) (interface{}, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if _dataCenter.isExpired(key) {
		return nil, fmt.Errorf("not find " + key)
	}
	value, ok := this.data[key]
	var err error
	if ok {
		err = nil
	} else {
		err = fmt.Errorf("not find " + key)
	}
	return value.body, err
}

func (this *dataPoll) Set(key string, value interface{}) { //set 如果没有就会添加 如果有就会覆盖
	mutex.Lock()
	defer mutex.Unlock()
	this.data[key] = &dataCenterKeyInfo{body: value, key: key, addtime: time.Now().UnixNano(), gtime: -1} //没有过期时间的默认-1标识
}

func (this *dataPoll) SetKeyAndInfo(key string, value interface{}, g int64) { //此处注意要纳秒级别
	mutex.Lock()
	defer mutex.Unlock()
	nowtime := time.Now().UnixNano()
	this.data[key] = &dataCenterKeyInfo{body: value, key: key, addtime: nowtime, gtime: nowtime + g}
}

func (this *dataPoll) clearDataCenter() {
	mutex.Lock()
	defer mutex.Unlock()
	now_time := time.Now().UnixNano() // 不需要太精确只是说定期清理而已
	for k, v := range this.data {
		fmt.Println(k, v)
		if v.gtime <= now_time {
			delete(this.data, k)
		}
	}
}

func (this *dataPoll) isExpired(key string) bool { //过期或者是没有这个值
	data, ok := this.data[key]
	if !ok {
		log.Println("not " + key + "key")
		return true
	}
	return data.gtime > time.Now().UnixNano()
}

func (this *dataPoll) isKey(key string) bool { //过期或者是没有这个值
	_, ok := this.data[key]
	return ok
}
