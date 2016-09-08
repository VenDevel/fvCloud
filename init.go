package main

import (
	"fmt"
	"fvCloud/sqlite"
	"fvCloud/utils"
	"fvCloud/utils/dataCenter"
)

func init() {
	ReadfvCouldConf()
	sqlite.RegisterDB()

}

func ReadfvCouldConf() {
	body, err := utils.ReadFileAllByPath("./conf/fvCloud.conf")
	if err != nil {
		return
	}
	mp, err := utils.ByteJsonToMap(body)
	if err != nil {
		return
	}
	fmt.Println(mp)

	datacenter := dataCenter.GetInstance()
	datacenter.Add("fvConf", mp)
	fmt.Println(mp, datacenter)
}
