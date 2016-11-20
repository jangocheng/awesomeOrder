package utils

import (
	"fmt"
	"testing"
)

var m1 = make(map[string]interface{})


func TestParse(t *testing.T) {
	if _, err := ParseConfig("test.yaml"); err == nil {
		fmt.Println(err)
		return
	}
	if _,err:=ParseConfig("");err!=nil{
		fmt.Println(err)
		return
	}
	if _,err:=ParseConfig("../conf/awesome.yaml");err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(AwesomeConfig["redis"])
	ma:=AwesomeConfig["redis"].(CONF)
	fmt.Println(ma["addr"])
	fmt.Println(ma["pwd"])
	fmt.Println(ma["db"])
}

func TestMap(t *testing.T)  {
	m2 := returnMap()
	m2["watermellon"]=1
	fmt.Println("m1 after:",m1)

}


func returnMap() map[string]interface{} {
	m2 := map[string]interface{}{}

	m2["apple"]=1
	m2["pear"]=1

	for k,v:=range m2  {
		m1[k]=v
	}
	//m1 = m2
	fmt.Println("m1 before:",m1)
	return m2
}

