package utils

import (
	"io/ioutil"
	"os"

	"awesomeOrder/logs"

	"gopkg.in/yaml.v2"
	"fmt"
)

const yamlpath string = "../conf/awesome.yaml"
type CONF map[interface{}]interface{}
var AwesomeConfig CONF = make(CONF)

func ParseConfig(path string) (CONF,error) {
	if ""==path {
		path=yamlpath
	}
	file,err:=os.OpenFile(path,os.O_RDONLY,os.ModePerm)
	if err!=nil {
		logs.Logger.Errorf("open %v err:%v",path,err.Error())
		return nil,err
	}
	defer file.Close()
	b,err:=ioutil.ReadAll(file)
	if err!=nil {
		logs.Logger.Errorf("readfile %v err:%v",path,err.Error())
		return nil,err
	}
	config := CONF{}
	if err=yaml.Unmarshal(b,&config);err!=nil{
		logs.Logger.Errorf("unmarshal %s err:%v",string(b),err.Error())
		return nil,err
	}
	if yamlpath==path {
		AwesomeConfig = config
	}
	fmt.Println(config)
	return config,nil
}
