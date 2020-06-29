package json_demo

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	//ServerName string `json:"server_name"`
	//ServerIp string `json:"server_ip"`
	//ServerPort int `json:"server_port"`
	ServerName string
	ServerIp string
	ServerPort int
}


func SerializeStruct(){

	server := new(Server)
	server.ServerName = "服务器"
	server.ServerIp = "127.0.0.1"
	server.ServerPort = 80

	bt, err := json.Marshal(server)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return
	}

	fmt.Println("result = ", string(bt))

}

func SerializeMap()  {
	myMap := make(map[string]interface{})
	myMap["ServerName"] = "服务器"
	myMap["ServerIp"] = "127.0.0.1"
	myMap["ServerPort"] = 80
	//
	bt, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return
	}
	fmt.Println("result = ", string(bt))
	
}


func DeserializeStruct(){
	jsonStr := `{"ServerIp":"127.0.0.1","ServerName":"服务器","ServerPort":80}`
	server := new(Server)

	err := json.Unmarshal([]byte(jsonStr), &server)
	if err != nil{
		fmt.Println("err = ", err.Error())
		return
	}

	fmt.Println("DeserializeStruct server = ", server)
}

func DeserializeMap()  {
	jsonStr := `{"ServerIp":"127.0.0.1","ServerName":"服务器","ServerPort":80}`
	server := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &server)
	if err != nil{
		fmt.Println("err = ", err.Error())
		return
	}

	fmt.Println("DeserializeMap  server= ", server)
}