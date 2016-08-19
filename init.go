package main

import (
	"fmt"
	"fvCloud/utils"
)

func init() {
	ReadfvCouldConf()

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

	datacenter := utils.GetInstance()
	datacenter.Add("fvConf", mp)
	fmt.Println(mp, datacenter)
}
