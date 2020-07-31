package json

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string `json:"name,omitempty"`
	sex string  `json:"sex,omitempty"`
	age int8	`json:"age,omitempty"`
	mobile string `json:"iphone,omitempty"`
	address string `json:"address,omitempty"`
}

func SerialPerson()  {

	person := Person{
		name:    "jackie",
		sex:     "male",
		mobile:  "13520468276",
		address: "China",
	}

	if data, err :=json.Marshal(person);err ==nil{
		fmt.Printf("%s\n", data)
		fmt.Println("data = ", string(data))
	}else {
		fmt.Println("SerialPerson err!")
	}

}

func Unserial()  {
	//
	var person Person
	jsonStr := ``
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil{

	}

}
