package main

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"io/ioutil"
)

type FieldModel struct {
	accept_client_id []string
	field []string
	filter FilterModel
}

var New = func() *FieldModel {
	return &FieldModel{
		accept_client_id: nil,
		field:            nil,
		filter:           FilterModel{},
	}
}


type FilterModel struct {
	relation string
	condition []ConditionModel
}

type ConditionModel struct {
	key string
	opertor string
	value string
}


func main() {

	//res1 := gojsonq.New().FromString("[19, 90.9, 7, 67.5]").Sum()
	//fmt.Printf("%#v\n", res1)


	//all := gojsonq.New().File("/Users/yihangzhang/go/src/go-daydaystudy/jsonq.json").Get()
	//fmt.Println(all)

	//res := gojsonq.New().File("./jsonq.json").From("data.releated_info").Get()
	//fmt.Printf("%#v\n", res)


	buf, err := ioutil.ReadFile("./conf.json")
	fieldConf := FieldModel{}
	if err != nil {
		fmt.Println(err)
	}
	err1 := json.Unmarshal(buf, &fieldConf)
	if err1 != nil{
		fmt.Println(err)
	}
	fmt.Println(fieldConf)




	res1 := gojsonq.New().File("./jsonq.json")
	var aa interface{}
	json.Unmarshal([]byte(res1.String()), &aa)
	//fmt.Println(res1)

	res2 := gojsonq.New().File("./jsonq.json").From("data.releated_info").Select("app_id","area_code").Get()
	data, _ := json.Marshal(res2)
	fmt.Printf(string(data))

	//fmt.Printf("%#v\n", res2.Pluck("app_id"))

	//fmt.Printf("%#v\n", res2.Only("app_id","hosp_id"))

	//
	//fmt.Println(res2.Select("app_id","hosp_id").Select("dept_id").String())
	
}
